package types

import "github.com/google/uuid"

type Request struct {
	Server *uuid.UUID `json:"server,omitempty"`
	Token  string     `json:"token,omitempty"`
	Data   any        `json:"data"`
}

type Response struct {
	Server *uuid.UUID `json:"server,omitempty"`
	Data   any        `json:"data"`
}

type CreateRoomRequest struct {
}
