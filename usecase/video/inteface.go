package video

import (
	"context"

	"github.com/dailoi280702/vrs-general-service/type/model"
	"github.com/dailoi280702/vrs-general-service/type/request"
)

type I interface {
	GetVideoById(context.Context, request.GetById) (model.Video, error)
	GetVideoByIds(context.Context, request.GetByIds) ([]model.Video, error)
	GetVideoByUserHistory(context.Context, request.GetById) ([]model.Video, error)
	Update(context.Context, request.UpdateVideo) error
}
