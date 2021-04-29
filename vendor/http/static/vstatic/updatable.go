package vstatic

import (
	"encoding/json"
	"errors"
	"github.com/blang/semver"
	"http/static/ustatic"
	"net/http"
)

func (s *VersionStatic) SetUpdate(remote string) {
	s.internal.updater = ustatic.NewUStatic(remote, s.internal.dir)
}

func (s *internalVersionStatic) checkLatest() (*VersionInfo, error) {
	if s.updater == nil {
		return nil, errors.New("no remote update was set")
	}
	var latest, err = s.updater.GetLatest()
	if err != nil {
		return nil, err
	}
	versions, err := s.ListVersion()
	if err != nil {
		return nil, err
	}
	latestVersion, err := semver.ParseTolerant(latest.Version)
	if err != nil {
		return nil, err
	}
	for _, v := range versions {
		version, err := semver.ParseTolerant(v.Version)
		if err != nil {
			continue
		}
		if latestVersion.LTE(version) {
			return nil, nil
		}
	}
	return &VersionInfo{
		Version: latest.Version,
		Path:    latest.Path,
		Active:  s.active == latest.Version,
	}, nil
}

func (s *internalVersionStatic) sendJson(w http.ResponseWriter, data interface{}) {
	buffer, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buffer)
}

func (s *internalVersionStatic) handleCheckLatest(w http.ResponseWriter, r *http.Request) {
	var info, err = s.checkLatest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if info == nil {
		w.Write([]byte("already up-to-date"))
		return
	}
	s.sendJson(w, info)
}

func (s *internalVersionStatic) handleDownload(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var version = q.Get("version")
	var path = q.Get("path")
	var err = s.updater.Download(version, path)
	if err != nil {
		http.Error(w, "download failed "+err.Error(), http.StatusInternalServerError)
		return
	}
	s.sendJson(w, nil)
}

func (s *internalVersionStatic) handleDownloadLatest(w http.ResponseWriter, r *http.Request) {
	var info, err = s.checkLatest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if info == nil {
		w.Write([]byte("already update-to-date"))
		return
	}
	err = s.updater.Download(info.Version, info.Path)
	if err != nil {
		http.Error(w, "download failed "+err.Error(), http.StatusInternalServerError)
		return
	}
	s.sendJson(w, nil)
}

func (s *internalVersionStatic) handleUpdate(w http.ResponseWriter, r *http.Request) {
	var info, err = s.autoUpdate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.sendJson(w, info)
}

func (s *internalVersionStatic) autoUpdate() (*VersionInfo, error) {
	if s.updater == nil {
		return nil, errors.New("no update remote found")
	}
	var info, err = s.checkLatest()
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("already up-to-date")
	}
	err = s.updater.Download(info.Version, info.Path)
	if err != nil {
		logger.Errorf("download failed %s", err.Error())
		return nil, err
	}
	err = s.activate(info.Version)
	if err != nil {
		logger.Errorf("activate faield %s", err.Error())
		return nil, err
	}
	if err = s.saveVersion(); err != nil {
		logger.Errorf("save version %s", err.Error())
		return nil, err
	}
	return info, nil
}
