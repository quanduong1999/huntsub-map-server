package config

import (
	"fmt"
	"huntsub/huntsub-map-server/config/database"
	"huntsub/huntsub-map-server/config/shared"
	"huntsub/huntsub-map-server/config/station"
)

var logger = shared.ConfigLog

type ProjectConfig struct {
	// Business business.BusinessConfig `json:"business"`
	Database database.DatabaseConfig `json:"database"`
	Station  station.StationConfig   `json:"station"`
}

func (p ProjectConfig) String() string {
	return fmt.Sprintf("config:[%s][%s][%s]", p.Database, p.Station)
}

func (p *ProjectConfig) Check() {
	p.Station.Check()
	p.Database.Check()
	// p.Business.Check()
}
