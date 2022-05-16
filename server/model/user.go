package model

import "github.com/google/uuid"

// ユーザー
type User struct {
	ID     uuid.UUID // ID
	Name   string    // ユーザー名
	IconID uuid.UUID // アイコンID
}
