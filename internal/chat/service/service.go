package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Config struct {
	Port int
}

type ChatService struct {
	logger     *log.Logger
	authHelper *authHelper
	cfg        Config
}

func New(cfg Config) (*ChatService, error) {
	authHelper, err := newAuthHelper()
	if err != nil {
		return nil, err
	}
	return &ChatService{
		logger:     log.New(os.Stdout, "[Chat Server] ", log.LstdFlags),
		cfg:        cfg,
		authHelper: authHelper,
	}, nil
}

func (s *ChatService) Start() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedHeaders: []string{"*"},
		Debug:          true,
	}))

	r.Get("/ws", s.handleWs)
	r.Post("/authenticate", s.authHelper.Authenticate)
	r.Route("/room", func(r chi.Router) {
		r.Use(s.authHelper.WithAuthCtx)
		r.Post("/{roomId}", s.JoinRoom)
	})

	s.logger.Printf("Starting chat server on :%d", s.cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), r)

}
