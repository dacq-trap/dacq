package service

import (
	"context"

	"github.com/dacq-trap/dacq/server/model"
)

type UsersService interface {
	// ユーザーをNameで取得
	//
	// 存在しなければ、model.ErrNotFoundを返す
	ReadUserByName(ctx context.Context, name string) (user *model.User, err error)

	// ユーザーを作成
	CreateUser(ctx context.Context, name string) (user *model.User, err error)
}
