package httpserver

import (
	"http/socket"
	"huntsub/huntsub-map-server/config"
	"huntsub/huntsub-map-server/x/mlog"
	"net/http"

	"github.com/golang/glog"
)

var logger = mlog.NewTagLog("httpserver")

type ProjectHttpServer struct {
	pc      *config.ProjectConfig
	ready   chan struct{}
	handler http.Handler
	Ws      *socket.Hub
}

func NewProjectHttpServer(pc *config.ProjectConfig) *ProjectHttpServer {
	var s = &ProjectHttpServer{
		pc:    pc,
		ready: make(chan struct{}, 2),
		Ws:    socket.NewHub(),
	}
	s.handler = s.makeHandler()
	return s
}

func (s *ProjectHttpServer) listen() {
	var addr = s.pc.Station.Server.Addr()
	var server = http.Server{
		Addr:         addr,
		TLSNextProto: nil,
		Handler:      s.handler,
	}
	logger.Infof(0, "Listening on http://%s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		glog.Errorf("Server %s", err.Error())
		glog.Flush()
	}
}

func (s *ProjectHttpServer) listenTLS() {
	serverConfig := s.pc.Station.Server
	if serverConfig.HasHttps() {
		if err := serverConfig.Check(); err != nil {
			logger.Errorf("Cannot start HTTPS server due to [%s]", err.Error())
			return
		}
		var addr = serverConfig.AddrTLS()
		var server = http.Server{
			Addr:    addr,
			Handler: s.handler,
		}
		logger.Infof(0, "Listening on https://%s\n", addr)
		if err := server.ListenAndServeTLS(serverConfig.Certificate, serverConfig.PrivateKey); err != nil {
			glog.Errorf("Server %s", err.Error())
		}
	}
}

func (s *ProjectHttpServer) shutdown() {
	logger.Infoln(0, "Shutting down the server...")
	// ctx, _ := s.config.Wait()
	// err := s.Shutdown(ctx)
	// if err == nil {
	// 	logger.Infoln(0, "Server gracefully stopped")
	// } else {
	// 	logger.Errorf("Server shutdown %s\n", err.Error())
	// }
}
