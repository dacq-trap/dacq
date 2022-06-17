package model

import (
	"time"

	"github.com/dacq-trap/dacq/server/utils/optional"
)

// サブミッション
type Submission struct {
	ID          int                 // ID
	Competition CompetitionCoreInfo // コンペティション
	Team        TeamCoreInfo        // チーム
	Submitter   User                // 提出者
	Comment     string              // コメント
	FileID      string              // 提出されたCSVファイルのID
	Score       optional.Float      // スコア
	SubmittedAt time.Time           // 提出日時
}
