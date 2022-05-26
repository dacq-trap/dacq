package model

import (
	"time"

	"github.com/google/uuid"
)

// スコアの順位決定順
type Order bool

var (
	Asc  Order = true  // スコアが高い順
	Desc Order = false // スコアが低い順
)

// コンテストの基本情報
type ContestCoreInfo struct {
	ID   uuid.UUID // ID
	Name string    // コンテスト名
}

// コンテスト
type Contest struct {
	ContestCoreInfo
	Author          User           // コンテスト作成者
	Rule            string         // ルール
	DataID          string         // 学習データファイルID
	DataDescription string         // 学習データの説明
	JudgementID     int            // スコア判定基準のID
	AnswerDataID    string         // 正解データファイルID
	Teams           []TeamCoreInfo // 参加チーム
	HeldFrom        time.Time      // コンテスト開始日時
	HeldUntil       time.Time      // コンテスト終了日時
}
