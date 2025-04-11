package auth

import (
	"github.com/dailoi280702/vrs-general-service/repository"
	"github.com/dailoi280702/vrs-general-service/repository/user"
)

type Usecase struct {
	UserRepo user.I
}

func New(repo *repository.Repository) I {
	return &Usecase{
		UserRepo: repo.UserRepo,
	}
}
