package models

type User struct {
	ID uint64 `json:"id,omitempty" gorm:"column:id;primarykey"`
	Name string `json:"name,omitempty" gorm:"unique;not null"` // Must be unique
	Nickname string `json:"nickname,omitempty" gorm:"not null"` // Not necessary unique 
	Password string `json:"password,omitempty" gorm:"not null"`
	// Torrents []Torrent `json:"torrents"`
}