package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	DistPath string
	Port     int
}

type Service struct {
	cfg    Config
	logger *log.Logger
}

func New(cfg Config) Service {
	return Service{
		cfg:    cfg,
		logger: log.New(os.Stdout, "[Chat Web Server] ", log.LstdFlags),
	}
}

func (s *Service) Start() {
	pageMux := http.NewServeMux()
	pageMux.Handle("/", http.FileServer(http.Dir(s.cfg.DistPath)))
	s.logger.Printf("Starting web server on :%d", s.cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), pageMux)
}
