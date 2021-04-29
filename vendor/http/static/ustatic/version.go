package ustatic

import (
	"errors"
	"github.com/BurntSushi/toml"
	"net/http"
	"time"
)

type VersionInfo struct {
	Version string `json:"version"`
	Path    string `json:"path"`
	Active  bool   `json:"active"`
	MTime   int64  `json:"mtime"`
}

func (u *UStatic) GetLatest() (*VersionInfo, error) {
	var c = http.Client{
		Timeout: 30 * time.Minute,
	}

	resp, err := c.Get(u.getLink("latest"))
	if err != nil {
		return nil, err
	}
	var latest = &VersionInfo{}
	defer resp.Body.Close()
	if resp.StatusCode > 200 {
		return nil, errors.New("remote host response with  " + resp.Status)
	}
	_, err = toml.DecodeReader(resp.Body, latest)
	if err != nil {
		return nil, err
	}
	return latest, nil
}
