package usecase

import (
	"github.com/dailoi280702/vrs-general-service/repository"
	"github.com/dailoi280702/vrs-general-service/usecase/auth"
	"github.com/dailoi280702/vrs-general-service/usecase/video"
)

type Usecase struct {
	Auth  auth.I
	Video video.I
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		Auth:  auth.New(repo),
		Video: video.New(repo),
	}
}
