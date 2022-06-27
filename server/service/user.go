package service

import "github.com/dacq-trap/dacq/server/model"

type UserService interface {
	// ユーザーをNameで取得
	//
	// 存在しなければ、model.ErrNotFoundを返す
	ReadUserByName(name string) (*model.User, error)

	// ユーザーを作成
	CreateUser(name string) (*model.User, error)
}
