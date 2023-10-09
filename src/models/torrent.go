package models

type Torrent struct {
	ID uint `json:"id,omitempty" gorm:"column:id;primarykey"`
	Title string `json:"title,omitempty" gorm:"not null"`
	Likes uint `json:"likes,omitempty"`
	Dislikes uint `json:"dislikes,omitempty"`
	URI string `json:"uri,omitempty" gorm:"not null"`
	Image string `json:"img,omitempty"`
	Category string `json:"category,omitempty"`
}
