package usecase

import "gorm.io/gorm"

type Usecase struct{}

func NewUsecase(db *gorm.DB) *Usecase {
	return &Usecase{}
}
