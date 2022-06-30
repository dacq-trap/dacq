package repository

import (
	"github.com/dacq-trap/dacq/server/model"
	"github.com/google/uuid"
)

type TeamsRepository interface {
	// コンペティションIDでチームを取得
	SelectTeamsByCompetitionID(competitionID uuid.UUID) ([]*model.Team, error)
}
