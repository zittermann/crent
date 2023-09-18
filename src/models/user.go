package models

type User struct {
	ID uint64
	Name string // Must be unique
	Nickname string // Not necessary unique 
	Password string
	Torrents []Torrent
}