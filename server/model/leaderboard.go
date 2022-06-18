package model

import "time"

// リーダーボード
type Leaderboard struct {
	Rank             int       // 順位
	Team             Team      // チーム
	BestScore        float64   // 最高スコア
	TotalSubmissions int       // のべ提出数
	LastSubmittedAd  time.Time // 最終提出日時
}
