package repositories

import (
	"errors"

	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByID(id uint) (u models.User, err error)
	FindByName(name string, p response.Page,
	) (matches []models.User, totalElements int64, err error)
	FindByNickname(nickname string, p response.Page,
	) (matches []models.User, totalElements int64, err error)
	Save(user models.User)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// FindByID implements IUserRepository.
func (t *UserRepository) FindByID(id uint,
) (userFound models.User, err error) {

	err = t.db.Find(&userFound).Error

	if err != nil {
		return userFound,
			errors.New("There is no User with that ID")
	}

	return userFound, nil

}

// FindByName implements IUserRepository.
func (t *UserRepository) FindByName(name string, p response.Page,
) (matches []models.User, totalElements int64, err error) {

	t.db.Where("name LIKE %?%", name).
		Limit(p.GetLimit()).
		Offset(p.GetOffset()).
		Find(&matches)

	t.db.Model(&models.User{}).
		Where("name LIKE %?%", name).
		Count(&totalElements)

	if len(matches) == 0 {
		return nil, 0,
			errors.New("There are no matches with that name")
	}

	return matches, totalElements, nil

}

// FindByNickname implements IUserRepository.
func (t *UserRepository) FindByNickname(nickname string, p response.Page,
) (matches []models.User, totalElements int64, err error) {

	t.db.Where("nickname LIKE %?%", nickname).
		Limit(p.GetLimit()).
		Offset(p.GetOffset()).
		Find(&matches)

	t.db.Model(&models.User{}).
		Where("nickname LIKE %?%", nickname).
		Count(&totalElements)

	if len(matches) == 0 {
		return nil, 0,
			errors.New("There are no matches with that nickname")
	}

	return matches, totalElements, nil

}

// Save implements IUserRepository.
func (t *UserRepository) Save(user models.User) {
	t.db.Create(&user)
}