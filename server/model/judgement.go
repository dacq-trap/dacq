package model

// スコア判定基準
type Judgement struct {
	ScoreScript   string // スコア計算スクリプト
	ModuleFile    string // 実行時に必要なモジュール管理ファイル
	PublicSetting string // public / private区分の設定ファイル
	ScoreOrder    Order  // スコアの順位決定順
}
