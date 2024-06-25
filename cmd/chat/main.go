package main

import chatSvc "chat/internal/chat/service"

func main() {
	chatSvc, err := chatSvc.New(chatSvc.Config{
		Port: 9090,
	})
	if err != nil {
		panic(err)
	}
	chatSvc.Start()
}
