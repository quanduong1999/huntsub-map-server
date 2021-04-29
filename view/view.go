package view

import (
	"fmt"
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/config/cons"
	"huntsub/huntsub-map-server/x/timer"
	"io/ioutil"
	"net/http"
	"strings"
)

type ViewServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewViewServer() *ViewServer {

	var s = &ViewServer{
		ServeMux: http.NewServeMux(),
	}
	var tpm = http.FileServer(http.Dir("static"))
	s.Handle("/", http.StripPrefix("/", tpm))
	s.HandleFunc("/s", s.HandleShortUrl)
	s.HandleFunc("/upload", s.HandleUpload)
	return s
}

func (s *ViewServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.Recover(w)
	header := w.Header()
	// header.Add("Access-Control-Allow-Origin", r.Header.Get("*"))
	// header.Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	// header.Add("Access-Control-Allow-Origin", r.Header.Get("Content-Type"))
	// header.Add("Access-Control-Allow-Origin", r.Header.Get("Accept"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,Huntsub-Token"
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Set("Access-Control-Allow-Headers", allowedHeaders)
	header.Set("Access-Control-Expose-Headers", "Authorization")
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
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}

func (s *ViewServer) HandleShortUrl(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var newUrl = cons.REDIRECT_VIEW + id
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

func (s *ViewServer) HandleUpload(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustGet(r)
	var _type = r.URL.Query().Get("type")
	var u, err = cache_user.Get(sess.UserID)
	web.AssertNil(err)

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// A particular naming pattern
	// path := fmt.Sprint("static/upload/user/")
	var path string
	if _type == "img" {
		path = fmt.Sprintf("static/upload/user/%s/%s/img", timer.TimeToDay(u.MTime), sess.UserID)
	} else {
		path = fmt.Sprintf("static/upload/user/%s/%s/video", timer.TimeToDay(u.MTime), sess.UserID)
	}

	var str = strings.Split(handler.Filename, ".")
	p := "upload-*." + str[len(str)-1]
	tempFile, err := ioutil.TempFile(path, p)
	println(p)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	s.SendData(w, cons.DIR_DOWNLOAD_EXCEL+tempFile.Name())
}
