package service

import "github.com/dacq-trap/dacq/server/model"

type UserService interface {
	// ユーザーをIDで取得
	ReadUserByID(id string) (*model.User, error)
}
