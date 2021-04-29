package httpserver

import (
	"http/gziphandler"
	"http/static/upload"
	"huntsub/huntsub-map-server/api"
	"huntsub/huntsub-map-server/socket"
	"huntsub/huntsub-map-server/view"
	"net/http"
	"regexp"
)

func webAssetGzipHandler(handler http.Handler) http.Handler {
	gzip := gziphandler.GzipHandler(handler)
	assetRegex, _ := regexp.Compile(".(js|css)$")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if assetRegex.MatchString(r.URL.Path) {
			gzip.ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (phs *ProjectHttpServer) addStaticHandler(s *http.ServeMux) {
	// p := phs.pc
	// staticConfig := p.Station.Static

	// var app = vstatic.NewVersionStatic(staticConfig.AppFolder)

	// s.Handle("/", http.RedirectHandler("/app/", http.StatusFound))
	// s.Handle("/", http.StripPrefix("/", webAssetGzipHandler(app)))

	var up = upload.NewUploadFileServer("static", 40960000)
	s.Handle("/static/", http.StripPrefix("/static/", up))
}

func (phs *ProjectHttpServer) makeHandler() http.Handler {
	var server = http.NewServeMux()
	// anhht
	phs.addStaticHandler(server)
	server.Handle("/", NewServerStatic(""))
	apiServer := api.NewApiServer()
	server.Handle("/api/",
		gziphandler.GzipHandler(http.StripPrefix("/api", apiServer)),
	)
	socketServer := socket.NewSocketServer(phs.Ws)
	server.Handle("/ws/", http.StripPrefix("/ws", socketServer))
	viewServer := view.NewViewServer()
	server.Handle("/v/", http.StripPrefix("/v", viewServer))
	go func() {
		phs.ready <- struct{}{}
	}()
	return server
}
