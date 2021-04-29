package station

import (
	"strings"
)

type ProxyConfig struct {
	Upload string
	Update string
	Record string
}

func (p *ProxyConfig) isNone(s string) bool {
	return strings.TrimSpace(strings.ToLower(s)) == "none"
}

func (p *ProxyConfig) NoneUpload() bool {
	return p.isNone(p.Upload)
}

func (p *ProxyConfig) NoneUpdate() bool {
	return p.isNone(p.Update)
}

func (p *ProxyConfig) NoneRecord() bool {
	return p.isNone(p.Record)
}
