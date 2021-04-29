package vstatic

import (
	"errors"
	"http/static"
	"http/static/ustatic"
	"huntsub/huntsub-map-server/x/mlog"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/blang/semver"
)

var logger = mlog.NewTagLog("vstatic")

func (v *internalVersionStatic) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("internal ping active=" + v.active))
}

func (v *internalVersionStatic) handleListVersion(w http.ResponseWriter, r *http.Request) {
	versions, err := v.ListVersion()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	v.sendJson(w, versions)
}

func (v *internalVersionStatic) handleActivate(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var version = q.Get("version")
	err := v.activate(version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = v.saveVersion()
	if err != nil {
		http.Error(w, "save version failed "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("activate version " + version + " successfull"))
}

type VersionInfo ustatic.VersionInfo

func isValidSemver(name string) bool {
	_, err := semver.ParseTolerant(name)
	return err == nil
}

func (s *internalVersionStatic) ListVersion() ([]*VersionInfo, error) {
	var versions = []*VersionInfo{}
	if s == nil {
		return versions, nil
	}
	var dirs, err = ioutil.ReadDir(s.dir)
	if err != nil {
		return versions, err
	}
	for _, d := range dirs {
		if d.IsDir() && isValidSemver(d.Name()) {
			versions = append(versions, &VersionInfo{
				Version: d.Name(),
				MTime:   d.ModTime().Unix(),
				Active:  s.active == d.Name(),
			})
		}
	}

	return versions, nil
}

const versionFile = ".version"

func (v *internalVersionStatic) readVersion() (*VersionInfo, error) {
	var info = &VersionInfo{}
	var file = filepath.Join(v.dir, versionFile)
	var _, err = os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Infof(0, "version file %v missing", file)
		} else {
			logger.Infof(0, "read version file %v failed %v", file, err.Error())
		}
		return nil, nil
	}
	_, err = toml.DecodeFile(file, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (v *internalVersionStatic) saveVersion() error {
	var info = &VersionInfo{Version: v.active}
	var file = filepath.Join(v.dir, versionFile)
	var writer, err = os.Create(file)
	if err != nil {
		return err
	}
	defer writer.Close()
	err = toml.NewEncoder(writer).Encode(info)
	return err
}

func (v *internalVersionStatic) activate(active string) error {
	folder := filepath.Join(v.dir, active)
	stat, err := os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("version " + active + " not found")
		}
		return err
	}
	if !stat.IsDir() {
		return errors.New("path " + active + " is not a folder")
	}

	v.active = active
	v.current = static.NewStatic(folder)
	return nil
}

func (v *internalVersionStatic) reactivate() {
	var info, err = v.readVersion()
	if err != nil {
		v.current = nil
		logger.Errorf("read version failed failed %v", err.Error())
		return
	}
	if info == nil {
		v.activate("")
	} else {
		v.activate(info.Version)
	}
	return
}
