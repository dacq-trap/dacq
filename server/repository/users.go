package repository

import (
	"context"

	"github.com/dacq-trap/dacq/server/model"
)

type UsersRepository interface {
	// ユーザーをNameで取得
	//
	// 存在しなければ、model.ErrNotFoundを返す
	SelectUserByName(ctx context.Context, name string) (user *model.User, err error)

	// ユーザーを作成
	CreateUser(ctx context.Context, user model.User) (err error)
}
