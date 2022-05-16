package model

import "github.com/google/uuid"

type Team struct {
	ID        uuid.UUID
	ContestID uuid.UUID
	Users     []User
	Name      string
}
