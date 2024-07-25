.PHONY: chat-server
chat-server: 
	go run ./cmd/chat/main.go

.PHONY: chat-server-dev
chat-server-dev:
	reflex -s -r '[.]go$\' go run ./cmd/chat/main.go

.PHONY: chat-client
chat-client: 
	go run ./cmd/chat-client/main.go

.PHONY: chat-web
chat-web: 
	go run ./cmd/chat-web-server/main.go

.PHONY: tools
tools:
	cd ./tools && go install github.com/cespare/reflex
