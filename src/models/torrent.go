package models

type Torrent struct {
	ID uint
	Title string
	Likes uint
	Dislikes uint
	URI string
	Image string
	Category string
}
