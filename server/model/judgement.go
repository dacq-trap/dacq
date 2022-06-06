package model

// スコアの順位決定順
type Order string

var (
	Asc  Order = "asc"  // スコアが高い順
	Desc Order = "desc" // スコアが低い順
)

// スコア判定基準
type Judgement struct {
	ID             int    // ID
	Name           string // 判定方法の名前
	RawScoreScript string // スコア計算スクリプト
	RawModuleFile  string // 実行時に必要なモジュール管理ファイル
	PublicSetting  []bool // public / private区分の設定ファイル
	ScoreOrder     Order  // スコアの順位決定順
}
