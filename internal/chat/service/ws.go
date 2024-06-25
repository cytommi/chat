package service

import (
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Message struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
}

func (s *ChatService) serveWs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		s.logger.Printf("error accepting websocket upgrade: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer c.CloseNow()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// spawn a go routine to just keep reading payloads
	msgs := make(chan Message, 100)
	go func() {
		for ctx.Err() == nil {
			var msg Message
			wsjson.Read(ctx, c, &msg)
			msgs <- msg
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
				s.logger.Printf("GOT MSG: [%s] %s", msg.Username, msg.Msg)

			case <-ticker.C:
				err := wsjson.Write(ctx, c, "ayo")
				if err != nil {
					s.logger.Printf("err sending msg: %s, exiting", err)
					return
				}
			}
		}
	}()

	<-done
}
