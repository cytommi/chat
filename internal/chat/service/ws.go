package service

import (
	"fmt"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type PayloadType string

const AuthPayload PayloadType = "AUTH"
const MsgPayload PayloadType = "MSG"

type Payload struct {
	Type string `json:"type"`
	Auth string `json:"auth"` // should be a jwt or smth
	Msg  string `json:"msg"`
}

func (s *ChatService) auth(username string) string {
	return username
}

func (s *ChatService) handleWs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		s.logger.Printf("error accepting websocket upgrade: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.CloseNow()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// spawn a go routine to just keep reading payloads
	msgs := make(chan Payload, 100)
	go func() {
		for ctx.Err() == nil {
			var p Payload
			wsjson.Read(ctx, conn, &p)

			switch PayloadType(p.Type) {
			case AuthPayload:
				fmt.Println("auth")

			case MsgPayload:
				if len(p.Msg) > 0 {
					fmt.Printf("msg: %s", p.Msg)
					msgs <- p
				}
			}
		}
	}()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				s.logger.Println("connection ctx done")
				return

			case msg := <-msgs:
				s.logger.Printf("GOT MSG: [%s] %s", "USER", msg.Msg)

				// case <-ticker.C:
				// 	err := wsjson.Write(ctx, conn, "ayo")
				// 	if err != nil {
				// 		s.logger.Printf("err sending msg: %s, exiting", err)
				// 		return
				// 	}
			}
		}
	}()
	<-done
}
