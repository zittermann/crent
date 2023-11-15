package services

import (
	"errors"

	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/repositories"
	"github.com/obskur123/crent/src/utils"
)

type IUserService interface {
	FindByID(id uint) (*models.User, error)
	FindByName(name string, page int, limit int) (response.Page, error)
	FindByNickname(nickname string, page int, limit int,
		) (p response.Page, err error)
	Save(user *models.User) error
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{
		repository: repository,
	}
}

// FindByID implements IUserService.
func (t *UserService) FindByID(id uint) (user *models.User, err error) {
	
	user = t.repository.FindByID(id)

	if user.ID == 0 {
		return nil, errors.New("There is no User with that ID")
	}

	return user, nil

}

// FindByName implements IUserService.
func (t *UserService) FindByName(name string, page int, limit int,
	) (p response.Page, err error) {

	p.Page = page
	p.Limit = limit

	matches, totalElements := t.repository.FindByName(name, p)

	if totalElements == 0 {
		return p, 
			errors.New("There are no matches with that name")
	}

	p.Content = *matches
	p.TotalElements = totalElements

	return p, nil

}

// FindByNickname implements IUserService.
func (t *UserService) FindByNickname(nickname string, page int, limit int,
	) (p response.Page, err error) {
	
		p.Page = page
		p.Limit = limit

		matches, totalElements := t.repository.FindByNickname(nickname, p)	

		if totalElements == 0 {
			return p, errors.New("There are no matches with that nickname")
		}
		
		p.Content = *matches
		p.TotalElements = totalElements

		return p, nil

}

// Save implements IUserService.
func (t *UserService) Save(user *models.User) error {

	if utils.IsEmpty(user.Name) {
		return errors.New("Name cannot empty string")
	}

	if utils.IsEmpty(user.Nickname) {
		return errors.New("Nickname cannot empty string")
	}

	if utils.IsEmpty(user.Password) {
		return errors.New("Password cannot empty string")
	}

	t.repository.Save(user)
	return nil
	
}
