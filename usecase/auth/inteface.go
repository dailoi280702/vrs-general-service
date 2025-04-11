package auth

import (
	"context"

	"github.com/dailoi280702/vrs-general-service/type/request"
	"github.com/dailoi280702/vrs-general-service/type/response"
)

type I interface {
	Register(ctx context.Context, req request.Register) (response.Register, error)
}
