package repository

import (
	"gorm.io/gorm"
)

type Repository struct{}

func New(getClient func() *gorm.DB) *Repository {
	return &Repository{}
}
