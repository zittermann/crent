package tests

import (
	"testing"

	"github.com/obskur123/crent/src/models"
	"github.com/obskur123/crent/src/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyTitle(t *testing.T) {
	testService := services.NewTorrentService(nil)

	torrent := models.Torrent {
		Title: "",
	}

	err := testService.Save(&torrent)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Title cannot empty string")

}

func TestValidateEmptyURI(t *testing.T) {
	testService := services.NewTorrentService(nil)

	torrent := models.Torrent {
		Title: "asd",
		URI: "",
	}

	err := testService.Save(&torrent)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "URI cannot empty string")

}
