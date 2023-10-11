package repositories

import (
	response "github.com/obskur123/crent/src/data/responses"
	"github.com/obskur123/crent/src/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByID(id uint) (u *models.User)
	FindByName(name string, p *response.Page,
	) (matches *[]models.User, totalElements int64)
	FindByNickname(nickname string, p *response.Page,
		) (matches *[]models.User, totalElements int64)
	Save(user *models.User)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// FindByID implements IUserRepository.
func (t *UserRepository) FindByID(id uint) (userFound *models.User) {

	t.db.Find(userFound)

	return userFound

}

// FindByName implements IUserRepository.
func (t *UserRepository) FindByName(name string, p *response.Page,
) (matches *[]models.User, totalElements int64) {

	t.db.Where("name LIKE %?%", name).
		Limit(p.GetLimit()).
		Offset(p.GetOffset()).
		Find(matches)

	t.db.Model(&models.User{}).
		Where("name LIKE %?%", name).
		Count(&totalElements)


	return matches, totalElements

}

// FindByNickname implements IUserRepository.
func (t *UserRepository) FindByNickname(nickname string, p *response.Page,
) (matches *[]models.User, totalElements int64) {

	t.db.Where("nickname LIKE %?%", nickname).
		Limit(p.GetLimit()).
		Offset(p.GetOffset()).
		Find(matches)

	t.db.Model(&models.User{}).
		Where("nickname LIKE %?%", nickname).
		Count(&totalElements)

	return matches, totalElements

}

// Save implements IUserRepository.
func (t *UserRepository) Save(user *models.User) {
	t.db.Create(user)
}