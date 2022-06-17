package model

import (
	"time"

	"github.com/dacq-trap/dacq/server/utils/optional"
	"github.com/google/uuid"
)

// コンペティションの基本情報
type CompetitionCoreInfo struct {
	ID      uuid.UUID // ID
	Name    string    // コンペティション名
	StartAt time.Time // コンペティション開始日時
	EndAt   time.Time // コンペティション終了日時
}

// コンペティション
type Competition struct {
	CompetitionCoreInfo
	Author          User            // コンペティション作成者
	Rule            string          // ルール
	DataID          optional.String // 学習データファイルID
	DataDescription string          // 学習データの説明
	JudgementID     int             // スコア判定基準のID
	PublicSetting   map[int]bool    // public / private区分の設定
	AnswerDataID    optional.String // 正解データファイルID
	Teams           []TeamCoreInfo  // 参加チーム
}
