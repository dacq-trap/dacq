package repository

import "github.com/dacq-trap/dacq/server/model"

type UserRepository interface {
	// ユーザーをIDで取得
	SelectUserByID(id string) (*model.User, error)
}
