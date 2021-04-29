package station

import (
	"fmt"
	"huntsub/huntsub-map-server/config/shared"
	"huntsub/huntsub-map-server/x/conf"
)

var logger = shared.ConfigLog

type StationConfig struct {
	// Sync    SyncConfig
	// Storage StorageConfig
	Static StaticConfig
	Log    conf.LogConfig
	Server conf.ServerConfig
	// Sys     SystemConfig
	// Proxy   ProxyConfig
}

func (s *StationConfig) Check() {
	s.Log.Init()
	// s.Sync.Check()
	// s.Storage.Check()
	s.Static.Check()
	// s.Sys.check()
}

func (s StationConfig) String() string {
	return fmt.Sprintf("station:[%s][%s][%s]", s.Static, s.Log, s.Server)
	// return fmt.Sprintf("station:[%s][%s][%s][%s][%s]", s.Sync, s.Storage, s.Static, s.Log, s.Server)
}

// func (s *StationConfig) GetUploadLink() *url.URL {
// 	if s.Proxy.NoneUpload() || s.Sync.IsMaster() {
// 		return nil
// 	}
// 	m := s.Sync.GetMaster()
// 	return m
// }

// func (s *StationConfig) GetUpdateLink() *url.URL {
// 	if s.Proxy.NoneUpdate() || s.Sync.IsMaster() {
// 		return nil
// 	}
// 	m := s.Sync.GetMaster()
// 	return m
// }
