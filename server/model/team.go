package model

import (
	"github.com/google/uuid"
)

// チーム
type Team struct {
	ID           uuid.UUID    // ID
	ContestID    uuid.UUID    // コンテストID
	Users        []User       // 所属ユーザー
	Name         string       // チーム名
	IsMerging    bool         // マージ中かどうか(マージ中はできない処理を設けるため)
	Submissions  []Submission // チームによる全てのサブミッション
	HighestScore int          // チームの最高スコア
}
