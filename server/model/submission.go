package model

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

// サブミッション
type Submission struct {
	ID          uuid.UUID       // ID
	Contest     ContestCoreInfo // コンテスト
	Team        TeamCoreInfo    // チーム
	Submitter   User            // 提出者
	CodeURL     url.URL         // コードのURL
	Comment     string          // コメント
	Score       float64         // スコア
	SubmittedAt time.Time       // 提出日時
}
