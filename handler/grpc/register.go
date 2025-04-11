package grpc

import (
	"context"
	"errors"

	"github.com/dailoi280702/vrs-general-service/client/mysql"
	"github.com/dailoi280702/vrs-general-service/log"
	"github.com/dailoi280702/vrs-general-service/proto"
	"github.com/dailoi280702/vrs-general-service/repository"
	"github.com/dailoi280702/vrs-general-service/type/request"
	"github.com/dailoi280702/vrs-general-service/usecase"
	"github.com/dailoi280702/vrs-general-service/usecase/auth"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.User, error) {
	usecase := usecase.New(repository.New(mysql.GetClient))
	logger := log.Logger()

	user, err := usecase.Auth.Register(ctx, request.Register{
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		var validationErrrs validator.ValidationErrors

		if errors.As(err, &validationErrrs) {
			return nil, status.Error(codes.InvalidArgument, validationErrrs.Error())
		} else if errors.Is(err, auth.ErrEmailDuplicated) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		logger.ErrorContext(ctx, "Failed to register user", "error", err, "request", req)

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.User{
		Email:    user.Email,
		FullName: user.FullName,
	}, nil
}
