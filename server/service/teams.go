package service

import (
	"context"

	"github.com/dacq-trap/dacq/server/model"
	"github.com/google/uuid"
)

type TeamsService interface {
	// コンペティションIDでチームを取得
	ReadTeamsByCompetitionID(ctx context.Context, competitionID uuid.UUID) ([]*model.Team, error)

	// ユーザーネームでチームを取得
	ReadTeamsByUserName(ctx context.Context, userName string) ([]*model.Team, error)

	// ユーザーネームとコンペティションIDでチームを取得
	//
	// 存在しなければ、model.ErrNotFoundを返す
	ReadTeamByUserNameAndCompetitionID(ctx context.Context, userName string, competitionID uuid.UUID) (*model.Team, error)

	// IDでチームを取得
	//
	// 存在しなければ、model.ErrNotFoundを返す
	ReadTeamByID(ctx context.Context, teamID uuid.UUID) (*model.Team, error)
}
