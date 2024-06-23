package service

import (
	"log"
	"net/http"
	"os"
)

type ChatService struct {
	logger *log.Logger
}

func NewChatService() (*ChatService, error) {
	return &ChatService{
		logger: log.New(os.Stdout, "[Chat Server] ", log.LstdFlags),
	}, nil
}

func (s *ChatService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.serveWs)

	http.ListenAndServe(":8000", mux)
}
