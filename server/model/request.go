package model

// チームマージリクエスト
type MergeRequest struct {
	From Team // マージ元チーム
	To   Team // マージ先チーム
}
