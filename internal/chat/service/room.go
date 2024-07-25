package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JwtIssuer        = "chat"
	JwtSubject       = "chat-room-token"
	JwtSigningMethod = jwt.SigningMethodRS256
)

type RoomClaims struct {
	RoomId string
	UserId string
	jwt.RegisteredClaims
}

var rooms = map[string][]string{}

type JoinRoomRequest struct {
	UserId string `json:"userId"`
}

func (s *ChatService) JoinRoom(w http.ResponseWriter, r *http.Request) {
	userId, err := UserFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	s.logger.Printf("User %s is joining room", userId)

	roomId := chi.URLParam(r, "roomId")
	if _, ok := rooms[roomId]; !ok {
		rooms[roomId] = []string{}
	}

	rooms[roomId] = append(rooms[roomId], userId)
}
