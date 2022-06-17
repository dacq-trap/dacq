package model

import (
	"github.com/google/uuid"
)

// チームの基本情報
type TeamCoreInfo struct {
	ID   uuid.UUID // ID
	Name string    // チーム名
}

// チーム
type Team struct {
	TeamCoreInfo
	Competition CompetitionCoreInfo // コンペティション
	Users       []User              // 所属ユーザー
	IsMerging   bool                // マージ中かどうか(マージ中はできない処理を設けるため)
}
