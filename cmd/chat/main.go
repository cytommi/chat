package main

import "chat/internal/chat/service"

func main() {
	chatSvc, err := service.NewChatService()
	if err != nil {
		panic(err)
	}
	chatSvc.Start()
}
