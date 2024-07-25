package room

import (
	"errors"
	"github.com/google/uuid"
)

var ErrAlreadyJoinedRoom = errors.New("member already joined room")

type Member struct {
	Id          uuid.UUID
	displayName string
}

type Room struct {
	Id      uuid.UUID
	Members map[uuid.UUID]Member
}

func New(id uuid.UUID, members []Member) *Room {
	mems := make(map[uuid.UUID]Member)
	for _, m := range members {
		mems[m.Id] = m
	}

	return &Room{
		Id:      id,
		Members: mems,
	}
}

func (r *Room) Join(newMember Member) error {
	if _, ok := r.Members[newMember.Id]; ok {
		return ErrAlreadyJoinedRoom
	}
	r.Members[newMember.Id] = newMember
	return nil
}
