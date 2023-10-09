package repositories

import (
	"errors"

	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/models"
	"gorm.io/gorm"
)

type ITorrentRepository interface {
	FindByID(id uint) (t models.Torrent, err error)
	FindByTitle(title string, p response.Page,
		) (torrents []models.Torrent, totalElements int64, err error)
	Save(t models.Torrent)
}

type TorrentRepository struct {
	db *gorm.DB
}

func NewTorrentRepository(db *gorm.DB) ITorrentRepository {
	return &TorrentRepository{db: db}
}

// FindByID implements ITorrentRepository.
func (t *TorrentRepository) FindByID(id uint,
	) (torrentFound models.Torrent, err error) {
	
	err = t.db.Find(&torrentFound, id).Error

	if err != nil {
		return torrentFound, 
			errors.New("Does not exists a torrent with that ID")
	}

	return torrentFound, nil

}

// FindByTitle implements ITorrentRepository.
func (t *TorrentRepository) FindByTitle(title string, p response.Page,
	) (torrents []models.Torrent, totalElements int64, err error) {

	t.db.
		Where("title LIKE %?%", title).
		Limit(p.GetLimit()).
		Offset(p.GetOffset()).
		Find(&torrents)

	t.db.Model(&models.Torrent{}).
		Where("title LIKE %?%", title).
		Count(&totalElements)

	if len(torrents) == 0 {
		return nil, 0, 
			errors.New("There are no matches with that title")
	}

	return torrents, totalElements, nil 

}

// Save implements ITorrentRepository.
func (t *TorrentRepository) Save(torrent models.Torrent) {
	t.db.Create(&torrent)
}
