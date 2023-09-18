package models

type Torrent struct {
	ID uint64
	Title string
	Likes uint64
	Dislikes uint64
	URI string
	Image string
	Category string
}
