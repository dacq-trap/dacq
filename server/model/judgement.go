package model

// スコアの順位決定順
type Order bool

var (
	Asc  Order = true  // スコアが高い順
	Desc Order = false // スコアが低い順
)

// スコア判定基準
type Judgement struct {
	RawScoreScript string // スコア計算スクリプト
	RawModuleFile  string // 実行時に必要なモジュール管理ファイル
	PublicSetting  []bool // public / private区分の設定ファイル
	ScoreOrder     Order  // スコアの順位決定順
}
