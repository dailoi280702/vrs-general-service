package repository

import (
	"github.com/dailoi280702/vrs-general-service/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo user.I
}

func New(getDB func() *gorm.DB) *Repository {
	return &Repository{
		UserRepo: user.NewRepo(getDB),
	}
}
