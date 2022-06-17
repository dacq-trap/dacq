package model

import (
	"time"

	"github.com/dacq-trap/dacq/server/utils/optional"
)

// リーダーボード
type Leaderboard struct {
	Rank             int            // 順位
	Team             Team           // チーム
	BestScore        optional.Float // 最高スコア
	TotalSubmissions int            // のべ提出数
	LastSubmittedAd  time.Time      // 最終提出日時
}
