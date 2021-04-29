package upload

import (
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const maxFileNameLen = 128

func makeSafeFilename(name string) string {
	v := slugify(name, true)
	if len(v) > maxFileNameLen {
		return v[:maxFileNameLen]
	}
	return v
}

func (s *UploadFileServer) postFile(w http.ResponseWriter, r *http.Request) {
	name := makeSafeFilename(s.filename(r))
	filename := filepath.Join(s.folder, s.requestFolder(r), name)
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte("fail to read file form " + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	outstream, err := os.Create(filename)
	if err != nil {
		glog.Error("create ", filename, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer outstream.Close()
	instream := &io.LimitedReader{N: s.maxFileSize, R: file}
	_, err = io.Copy(outstream, instream)
	if err != nil {
		glog.Error("save", filename, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(okString))
	w.WriteHeader(http.StatusOK)
}

const uploadForm = `
	<html>
		<head>
			<title>UploadFile</title>
		</head>
		<body>
			<form enctype="multipart/form-data" action="./" method="post">
				<input type="string" name="name" placeholder="filename" />
				<br>
				<input type="string" name="folder" placeholder="folder" />
				<br>
				<input type="file" name="file" />
				<br>
				<input type="submit" value="upload" />
			</form>
			<div>
				<a href=".."> back </a>
			</div>
		</body>
	</html>
`

const okString = `
	<html>
		<body>
			ok
			<a href="."> refresh </a>
		</body>
	</html>
`

func replyUploadForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(uploadForm))
}
