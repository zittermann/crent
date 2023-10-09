package models

import "gopkg.in/guregu/null.v4"

type Torrent struct {
	ID uint `json:"id,omitempty" gorm:"column:id;primarykey"`
	Title null.String `json:"title,omitempty" gorm:"not null"`
	Likes uint `json:"likes,omitempty"`
	Dislikes uint `json:"dislikes,omitempty"`
	URI string `json:"uri,omitempty" gorm:"not null"`
	Image null.String `json:"img,omitempty"`
	Category null.String `json:"category,omitempty"`
}
