package video

import (
	"github.com/dailoi280702/vrs-general-service/repository"
	"github.com/dailoi280702/vrs-general-service/repository/video"
)

type Usecase struct {
	VideoRepo video.I
}

func New(repo *repository.Repository) I {
	return &Usecase{
		VideoRepo: repo.VideoRepo,
	}
}
