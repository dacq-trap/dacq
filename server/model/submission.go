package model

import (
	"time"
)

// サブミッション
type Submission struct {
	ID          int                 // ID
	Contest     CompetitionCoreInfo // コンテスト
	Team        TeamCoreInfo        // チーム
	Submitter   User                // 提出者
	Comment     string              // コメント
	FileID      string              // 提出されたCSVファイルのID
	Score       float64             // スコア
	SubmittedAt time.Time           // 提出日時
}
