package model

import (
	"time"

	"github.com/google/uuid"
)

// コンテスト
type Contest struct {
	ID      uuid.UUID
	Name    string
	Author  uuid.UUID
	Problem Problem
	Teams   []Team
	HeldAt  time.Time
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
