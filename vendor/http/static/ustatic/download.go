package ustatic

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

func (u *UStatic) download(p string, dst string) error {
	var c = http.Client{
		Timeout: 30 * time.Minute,
	}
	resp, err := c.Get(u.getLink(p))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 200 {
		return errors.New("remote host response with  " + resp.Status)
	}
	tmpFile, err := ioutil.TempFile("", "download-"+p+"-")
	if err != nil {
		return err
	}
	// defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return err
	}
	tmpFile.Close()
	return unzip(tmpFile.Name(), dst)
}

func (u *UStatic) Download(version string, p string) error {
	var folder = path.Join(u.Local, version)
	err := os.MkdirAll(folder, os.ModeDir)
	if err != nil {
		return err
	}
	return u.download(p, folder)
}
