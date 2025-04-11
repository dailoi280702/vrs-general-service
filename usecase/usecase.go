package usecase

import (
	"github.com/dailoi280702/vrs-general-service/repository"
	"github.com/dailoi280702/vrs-general-service/usecase/auth"
)

type Usecase struct {
	Auth auth.I
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		Auth: auth.New(repo),
	}
}
