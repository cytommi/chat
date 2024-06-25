package main

import (
	chatSvc "chat/internal/web/service"
	"os"
	"path"
)

func main() {
	r, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	s := chatSvc.New(chatSvc.Config{
		Port:     8080,
		DistPath: path.Join(r, "build/clients/web/dist/"),
	})

	s.Start()
}
