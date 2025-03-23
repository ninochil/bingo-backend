package db

import "time"


// Host テーブル
type Host struct {
	HostID   string    `json:"host_id"`   // ホストID
	RoomCode string    `json:"room_code"` // 部屋のコード
	CreateAt time.Time `json:"create_at"` // 作成日時
	UpdateAt time.Time `json:"update_at"` // 更新日時
}

// Player テーブル
type Player struct {
	PlayerID string    `json:"player_id"` // 参加者ID
	HostID   string    `json:"host_id"`   // ホストID
	Name     string    `json:"name"`      // 参加者名
	CreateAt time.Time `json:"create_at"` // 作成日時
	UpdateAt time.Time `json:"update_at"` // 更新日時
}

// Question テーブル
type Question struct {
	QuestionID   string    `json:"question_id"`   // 質問ID
	Question     string    `json:"question"`      // 質問内容
	CreateAt     time.Time `json:"create_at"`     // 作成日時
	UpdateAt     time.Time `json:"update_at"`     // 更新日時
}

// QuestionUsage テーブル
type QuestionUsage struct {
	QuestionID   string    `json:"question_id"`   // 質問ID
	HostID       string    `json:"host_id"`       // ホストID
	PlayerID     string    `json:"player_id"`     // 参加者ID
	NumberOfUses int       `json:"number_of_uses"` // 使用回数
	CreateAt     time.Time `json:"create_at"`     // 作成日時
	UpdateAt     time.Time `json:"update_at"`     // 更新日時
}

// BingoCard テーブル
type BingoCard struct {
	BingoCardID string    `json:"bingo_card_id"` // ビンゴカードID
	HostID      string    `json:"host_id"`       // ホストID
	PlayerID    string    `json:"player_id"`     // 参加者ID
	IsBingo     bool      `json:"is_bingo"`      // ビンゴかどうか
	CreateAt    time.Time `json:"create_at"`     // 作成日時
	UpdateAt    time.Time `json:"update_at"`     // 更新日時
}

// BingoCardCellsStatus テーブル
type BingoCardCellsStatus struct {
	BingoCardID string    `json:"bingo_card_id"` // ビンゴカードID
	QuestionID  string    `json:"question_id"`   // 質問ID
	PositionX   int       `json:"position_x"`    // 横位置
	PositionY   int       `json:"position_y"`    // 縦位置
	IsChecked   *bool  `json:"isChecked,omitempty"`     // マスがチェックされているか
	CreateAt    time.Time `json:"create_at"`     // 作成日時
	UpdateAt    time.Time `json:"update_at"`     // 更新日時
}

