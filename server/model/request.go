package model

// チームマージリクエスト
type MergeRequest struct {
	From            TeamCoreInfo // マージ元チーム
	To              TeamCoreInfo // マージ先チーム
	AcceptanceQueue []User       // 承認待ちユーザー
}
