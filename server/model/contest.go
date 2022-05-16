package model

import (
	"time"

	"github.com/google/uuid"
)

// コンテストの基本情報
type ContestCoreInfo struct {
	ID   uuid.UUID // ID
	Name string    // コンテスト名
}

// コンテスト
type Contest struct {
	ContestCoreInfo
	Author  User           // コンテスト作成者
	Problem Problem        // 問題
	Teams   []TeamCoreInfo // チーム
	HeldAt  time.Time      // コンテスト開催日時
}

// スコアの順位決定順
type Order bool

var (
	Asc  Order = true  // スコアが高い順
	Desc Order = false // スコアが低い順
)

// 問題
type Problem struct {
	Rule          string   // ルール
	DataIDs       []string // 学習データファイルID
	ScoreScriptID string   // スコア計算スクリプトファイルID
	Order         Order    // スコアの順位決定順
}
