package model

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

// サブミッション
type Submission struct {
	ID          uuid.UUID // ID
	User        User      // 提出者
	FileID      string    // 提出されたファイルのID
	CodeURL     url.URL   // コードのURL
	Comment     string    // コメント
	Score       int       // スコア
	SubmittedAt time.Time // 提出日時
}
