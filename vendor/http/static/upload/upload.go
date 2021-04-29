package upload

import (
	"github.com/golang/glog"
	"http/static"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type UploadFileServer struct {
	folder        string
	maxFileSize   int64
	fileServer    http.Handler
	GetHandler    http.Handler
	PostHandler   http.Handler
	DeleteHandler http.Handler
}

const (
	oneGB = 1 << 30
	tenGB = 10 << 30
)

func NewUploadFileServer(folder string, maxFileSize int64) *UploadFileServer {
	if maxFileSize < oneGB {
		maxFileSize = oneGB
	} else if maxFileSize > tenGB {
		maxFileSize = tenGB
	}

	var s = &UploadFileServer{
		folder:      folder,
		fileServer:  static.NewStatic(folder),
		maxFileSize: maxFileSize,
	}
	s.GetHandler = http.HandlerFunc(s.getFile)
	s.PostHandler = http.HandlerFunc(s.postFile)
	s.DeleteHandler = http.HandlerFunc(s.deleteFile)
	return s
}

func (s *UploadFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/__") {
		if r.Method == http.MethodGet {
			replyUploadForm(w, r)
			return
		}
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/__")
	}
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
	header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	header.Add(
		"Access-Control-Max-Age",
		"2520000", // 30 days
	)

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		s.PostHandler.ServeHTTP(w, r)
	case http.MethodDelete:
		s.DeleteHandler.ServeHTTP(w, r)
	case http.MethodPut:
	case http.MethodGet:
		s.GetHandler.ServeHTTP(w, r)
	default:
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}
}

func (s *UploadFileServer) requestFolder(r *http.Request) string {
	var folder = r.FormValue("folder")
	if len(folder) < 2 {
		folder = path.Clean(r.URL.Path)
	}
	return folder
}

func (s *UploadFileServer) filename(r *http.Request) string {
	return r.FormValue("name")
}

func (s *UploadFileServer) deleteFile(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Join(s.folder, s.requestFolder(r), s.filename(r))
	err := os.Remove(filename)
	if err != nil {
		glog.Error("remove", filename, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *UploadFileServer) getFile(w http.ResponseWriter, r *http.Request) {
	s.fileServer.ServeHTTP(w, r)
}
