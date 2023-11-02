package repositories

import (
	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/models"
	"gorm.io/gorm"
)

type ITorrentRepository interface {
	FindByID(id uint) (*models.Torrent)
	FindByTitle(title string, p response.Page,
	) (*[]models.Torrent, int64)
	Save(t *models.Torrent)
}

type TorrentRepository struct {
	db *gorm.DB
}

func NewTorrentRepository(db *gorm.DB) ITorrentRepository {
	return &TorrentRepository{db: db}
}

// FindByID implements ITorrentRepository.
func (t *TorrentRepository) FindByID(id uint,
	) (*models.Torrent) {
	
	var torrentFound models.Torrent 

	t.db.Find(&torrentFound, id)

	return &torrentFound

}

// FindByTitle implements ITorrentRepository.
func (t *TorrentRepository) FindByTitle(title string, p response.Page,
	) (*[]models.Torrent, int64) {

	var torrents []models.Torrent
	var totalElements int64

	t.db.
		Where("title LIKE ?", title).
		Limit(p.GetLimit()).
		Offset(p.GetOffset()).
		Find(&torrents)

	t.db.Model(&models.Torrent{}).
		Where("title LIKE ?", title).
		Count(&totalElements)

	return &torrents, totalElements

}

// Save implements ITorrentRepository.
func (t *TorrentRepository) Save(torrent *models.Torrent) {
	t.db.Create(torrent)
}
