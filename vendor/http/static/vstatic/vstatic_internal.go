package vstatic

import (
	"http/static"
	"http/static/ustatic"
	"net/http"
	"time"
)

type internalVersionStatic struct {
	dir     string
	active  string
	mux     *http.ServeMux
	fs      http.Handler
	current http.Handler
	updater *ustatic.UStatic
}

func newInternalVersionStatic(dir string) *internalVersionStatic {
	var v = &internalVersionStatic{dir: dir}
	var mux = http.NewServeMux()
	v.mux = mux
	v.fs = static.NewStatic(dir)
	v.current = v.fs
	mux.HandleFunc("/list", v.handleListVersion)
	mux.HandleFunc("/activate", v.handleActivate)
	mux.HandleFunc("/check", v.handleCheckLatest)
	mux.HandleFunc("/download", v.handleDownload)
	mux.HandleFunc("/download_latest", v.handleDownloadLatest)
	mux.HandleFunc("/update", v.handleUpdate)
	mux.HandleFunc("/view", v.renderView)
	mux.Handle("/fs/", http.StripPrefix("/fs", v.fs))
	mux.HandleFunc("/", v.ping)
	return v
}

func (v *internalVersionStatic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.mux.ServeHTTP(w, r)
}

func (v *internalVersionStatic) eventLoop() {
	var tick = time.Tick(15 * time.Minute)
	for {
		select {
		case <-tick:
			var info, err = v.autoUpdate()
			if err != nil {
				logger.Errorf("update error %s", err.Error())
			} else {
				logger.Infof(0, "update to version %s doen", info.Version)
			}
		}
	}
}
