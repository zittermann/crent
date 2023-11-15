package test

import (
	"testing"

	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyName(t *testing.T) {
	testService := services.NewUserService(nil)

	user := models.User {
		Name: "",
	}

	err := testService.Save(&user)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Name cannot empty string")

}

func TestValidateEmptyNickname(t *testing.T) {
	testService := services.NewUserService(nil)

	user := models.User {
		Name: "test",
		Nickname: "",
	}

	err := testService.Save(&user)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Nickname cannot empty string")

}

func TestValidateEmptyPassword(t *testing.T) {
	testService := services.NewUserService(nil)

	user := models.User {
		Name: "test",
		Nickname: "test",
		Password: "",
	}

	err := testService.Save(&user)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Password cannot empty string")

}