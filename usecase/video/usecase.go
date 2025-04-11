package video

import (
	"context"
	"errors"

	spec "github.com/dailoi280702/vrs-general-service/repository/specification"
	"github.com/dailoi280702/vrs-general-service/type/model"
	"github.com/dailoi280702/vrs-general-service/type/request"
	"gorm.io/gorm"
)

var ErrVideoNotFound = errors.New("video not found")

var _ I = (*Usecase)(nil)

func (u *Usecase) Update(ctx context.Context, req request.UpdateVideo) error {
	if err := req.Validate(); err != nil {
		return err
	}

	video, err := u.VideoRepo.Get(ctx, spec.VideoByIds(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrVideoNotFound
		}

		return err
	}

	video.Comments = req.Comments
	video.Likes = req.Likes
	video.Views = req.Views
	video.Name = req.Name
	video.Shares = req.Shares
	video.WatchTime = req.WatchTime
	video.Length = req.Length

	return u.VideoRepo.Update(ctx, video)
}

func (u *Usecase) GetVideoById(ctx context.Context, req request.GetById) (model.Video, error) {
	var video model.Video
	var err error

	if err = req.Validate(); err != nil {
		return video, err
	}

	video, err = u.VideoRepo.Get(ctx, spec.VideoByIds(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return video, ErrVideoNotFound
		}
	}

	return video, err
}

func (u *Usecase) GetVideoByIds(ctx context.Context, req request.GetByIds) ([]model.Video, error) {
	return u.VideoRepo.Find(ctx, spec.VideoByIds(req.Ids...))
}

func (u *Usecase) GetVideoByUserHistory(ctx context.Context, req request.GetById) ([]model.Video, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return u.VideoRepo.Find(ctx, spec.VideoByUserHisotry(req.Id))
}
