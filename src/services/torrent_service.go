package services

import (
	"errors"

	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/repositories"
	"github.com/obskur123/crent/src/utils"
)

type ITorrentService interface {
	FindByID(id uint) (*models.Torrent, error)
	FindByTitle(title string, page int, limit int) (response.Page, error)
	Save(m *models.Torrent) error
}

type TorrentService struct {
	r repositories.ITorrentRepository
}

func NewTorrentService(
	repository repositories.ITorrentRepository,) ITorrentService {
	return &TorrentService{r: repository}
}

// FindByID implements ITorrentService.
func (t *TorrentService) FindByID(id uint) (torrent *models.Torrent, err error) {
	
	torrent = t.r.FindByID(id)

	if torrent.ID == 0 {
		return nil, errors.New("There is no Torrent with that ID")
	}

	return torrent, nil

}

// FindByTitle implements ITorrentService.
func (t *TorrentService) FindByTitle(title string,
	 page int, limit int) (p response.Page, err error) {
	
	p.Page = page
	p.Limit = limit

	torrentsList, totalElements := t.r.FindByTitle(title, p)

	if totalElements == 0 {
		return p, errors.New("There are no matches with that title")
	}

	p.Content = *torrentsList
	p.TotalElements = totalElements

	return p, nil

}

// Save implements ITorrentService.
func (t *TorrentService) Save(m *models.Torrent) error {

	if utils.IsEmpty(m.Title) {
		return errors.New("Title cannot empty string")
	}

	if utils.IsEmpty(m.URI) {
		return errors.New("URI cannot empty string")
	}

	t.r.Save(m)
	return nil
}
