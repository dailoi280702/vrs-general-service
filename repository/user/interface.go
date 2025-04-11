package user

import (
	"context"

	spec "github.com/dailoi280702/vrs-general-service/repository/specification"
	"github.com/dailoi280702/vrs-general-service/type/model"
)

type I interface {
	Count(ctx context.Context, spec spec.I) (int64, error)
	Get(ctx context.Context, spec spec.I) (model.User, error)
	Create(ctx context.Context, data model.User) error
}
