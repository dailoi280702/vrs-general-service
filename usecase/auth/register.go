package auth

import (
	"context"
	"errors"

	spec "github.com/dailoi280702/vrs-general-service/repository/specification"
	"github.com/dailoi280702/vrs-general-service/type/model"
	"github.com/dailoi280702/vrs-general-service/type/request"
	"github.com/dailoi280702/vrs-general-service/type/response"
	"github.com/dailoi280702/vrs-general-service/util"
)

var ErrEmailDuplicated = errors.New("email address already registered")

func (u *Usecase) Register(ctx context.Context, req request.Register) (response.Register, error) {
	resp := response.Register{}

	if err := req.Validate(); err != nil {
		return resp, err
	}

	count, err := u.UserRepo.Count(ctx, spec.UserByEmail(req.Email))
	if err != nil {
		return resp, err
	}

	if count > 0 {
		return resp, ErrEmailDuplicated
	}

	data := model.User{
		Email:    req.Email,
		FullName: req.FullName,
	}

	data.Password, err = util.HashPassword(req.Password)
	if err != nil {
		return resp, err
	}

	if err := u.UserRepo.Create(ctx, data); err != nil {
		return resp, err
	}

	resp.Email = req.Email
	resp.FullName = req.FullName

	return resp, nil
}
