package model

import (
	"net/url"
)

// ユーザー
type User struct {
	Name    string  // ユーザー名
	IconURL url.URL // アイコンURL
}
