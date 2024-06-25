package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port int
}

type ChatService struct {
	logger *log.Logger
	cfg    Config
}

func New(cfg Config) (*ChatService, error) {
	return &ChatService{
		logger: log.New(os.Stdout, "[Chat Server] ", log.LstdFlags),
		cfg:    cfg,
	}, nil
}

func (s *ChatService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.serveWs)

	s.logger.Printf("Starting chat server on :%d", s.cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), mux)
}
