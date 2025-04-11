package repository

import (
	"github.com/dailoi280702/vrs-general-service/repository/user"
	"github.com/dailoi280702/vrs-general-service/repository/video"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo  user.I
	VideoRepo video.I
}

func New(getDB func() *gorm.DB) *Repository {
	return &Repository{
		UserRepo:  user.NewRepo(getDB),
		VideoRepo: video.NewRepo(getDB),
	}
}
