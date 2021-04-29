package conf

import (
	"context"
	"fmt"
	"os"
	"time"
)

type ServerConfig struct {
	Host string
	Port int
	// HTTPS
	HTTPSPort   int
	Certificate string
	PrivateKey  string
	// Shutdown
	ShutdownWait int
}

func (s *ServerConfig) Addr() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

func (s *ServerConfig) AddrTLS() string {
	return fmt.Sprintf("%v:%v", s.Host, s.HTTPSPort)
}

func (s ServerConfig) String() string {
	return fmt.Sprintf("server:host=%v,port=%v", s.Host, s.Port)
}

func (s *ServerConfig) Wait() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(s.ShutdownWait)*time.Second)
}

func (s *ServerConfig) HasHttps() bool {
	return s.HTTPSPort > 1
}

func (s *ServerConfig) checkTLS() error {
	if _, err := os.Stat(s.Certificate); err != nil {
		return fmt.Errorf("Certificate %s error %s", s.Certificate, err.Error())
	}
	if _, err := os.Stat(s.PrivateKey); err != nil {
		return fmt.Errorf("Private Key %s error %s", s.PrivateKey, err.Error())
	}
	return nil
}

func (s *ServerConfig) Check() error {
	if s.HasHttps() {
		return s.checkTLS()
	}
	return nil
}
