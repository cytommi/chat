package service

import (
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func (s *ChatService) serveWs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		s.logger.Printf("error accepting websocket upgrade: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer c.CloseNow()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				s.logger.Println("connection ctx done")
				return
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
