package video

import (
	"context"

	spec "github.com/dailoi280702/vrs-general-service/repository/specification"
	"github.com/dailoi280702/vrs-general-service/type/model"
)

type I interface {
	Get(ctx context.Context, spec spec.I) (model.Video, error)
	Find(ctx context.Context, spec spec.I) ([]model.Video, error)
	Update(ctx context.Context, data model.Video) error
}
