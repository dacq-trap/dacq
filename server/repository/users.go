package repository

import "github.com/dacq-trap/dacq/server/model"

type UsersRepository interface {
	// ユーザーをNameで取得
	//
	// 存在しなければ、model.ErrNotFoundを返す
	SelectUserByName(name string) (*model.User, error)

	// ユーザーを作成
	CreateUser(model.User) error
}
