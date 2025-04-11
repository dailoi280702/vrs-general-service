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
	videouc "github.com/dailoi280702/vrs-general-service/usecase/video"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetUserWatchedHistory(ctx context.Context, req *proto.IdRequest) (*proto.Videos, error) {
	logger := log.Logger()
	uc := usecase.New(repository.New(mysql.GetClient))

	videos, err := uc.Video.GetVideoByUserHistory(ctx, request.GetById{Id: req.GetId()})
	if err != nil {
		var validationErrrs validator.ValidationErrors

		if errors.As(err, &validationErrrs) {
			return nil, status.Error(codes.InvalidArgument, validationErrrs.Error())
		}

		logger.ErrorContext(ctx, "Failed to get user watched history", "error", err, "request", req)

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &proto.Videos{}
	for _, v := range videos {
		res.Videos = append(res.Videos, &proto.Video{
			Id:        v.ID,
			Views:     v.Views,
			Likes:     v.Likes,
			Comments:  v.Comments,
			Shares:    v.Shares,
			Length:    v.Length,
			WatchTime: v.WatchTime,
		})
	}

	return res, nil
}

func (s *Service) GetVideosByIds(ctx context.Context, req *proto.GetVideosByIdsRequest) (*proto.Videos, error) {
	logger := log.Logger()
	uc := usecase.New(repository.New(mysql.GetClient))

	videos := req.GetId()
	resVideos, err := uc.Video.GetVideoByIds(ctx, request.GetByIds{Ids: videos})
	if err != nil {
		var validationErrrs validator.ValidationErrors

		if errors.As(err, &validationErrrs) {
			return nil, status.Error(codes.InvalidArgument, validationErrrs.Error())
		}

		logger.ErrorContext(ctx, "Failed to get videos by ids", "error", err, "request", req)

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &proto.Videos{}
	for _, v := range resVideos {
		res.Videos = append(res.Videos, &proto.Video{
			Id:        v.ID,
			Views:     v.Views,
			Likes:     v.Likes,
			Comments:  v.Comments,
			Shares:    v.Shares,
			Length:    v.Length,
			WatchTime: v.WatchTime,
		})
	}

	return res, nil
}

func (s *Service) GetVideoByID(ctx context.Context, req *proto.IdRequest) (*proto.Video, error) {
	logger := log.Logger()
	uc := usecase.New(repository.New(mysql.GetClient))

	video, err := uc.Video.GetVideoById(ctx, request.GetById{Id: req.GetId()})
	if err != nil {
		var validationErrrs validator.ValidationErrors

		if errors.As(err, &validationErrrs) {
			return nil, status.Error(codes.InvalidArgument, validationErrrs.Error())
		}

		if errors.Is(err, videouc.ErrVideoNotFound) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		logger.ErrorContext(ctx, "Failed to get video by id", "error", err, "request", req)

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &proto.Video{
		Id:        video.ID,
		Views:     video.Views,
		Likes:     video.Likes,
		Comments:  video.Comments,
		Shares:    video.Shares,
		Length:    video.Length,
		WatchTime: video.WatchTime,
	}

	return res, nil
}

func (s *Service) UpdateVideo(ctx context.Context, req *proto.Video) (*emptypb.Empty, error) {
	logger := log.Logger()
	uc := usecase.New(repository.New(mysql.GetClient))

	err := uc.Video.Update(ctx, request.UpdateVideo{
		Id:        req.GetId(),
		Likes:     req.GetLikes(),
		Views:     req.GetViews(),
		Comments:  req.GetComments(),
		Shares:    req.GetShares(),
		Length:    int64(req.GetLength()),
		WatchTime: int64(req.GetWatchTime()),
	})
	if err != nil {
		var validationErrrs validator.ValidationErrors

		if errors.As(err, &validationErrrs) {
			return nil, status.Error(codes.InvalidArgument, validationErrrs.Error())
		}

		if errors.Is(err, videouc.ErrVideoNotFound) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		logger.ErrorContext(ctx, "Failed to update video", "error", err, "request", req)

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
