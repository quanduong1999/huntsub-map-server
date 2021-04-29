package ustatic

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func unzipFile(dst string, zf *zip.File) error {
	var fileName = filepath.Join(dst, zf.Name)
	if zf.FileInfo().IsDir() {
		return os.MkdirAll(fileName, zf.Mode())
	}
	reader, err := zf.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, zf.Mode())
	if err != nil {
		return err
	}
	defer writer.Close()
	_, err = io.Copy(writer, reader)
	return err
}

func unzip(src, dst string) error {
	zfile, err := zip.OpenReader(src)
	if nil != err {
		return err
	}
	defer zfile.Close()

	//Create File
	for _, f := range zfile.File {
		err = unzipFile(dst, f)
		if err != nil {
			return err
		}
	}
	return nil
}
