package video_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dailoi280702/vrs-general-service/mock/video"
	"github.com/dailoi280702/vrs-general-service/type/model"
	"github.com/dailoi280702/vrs-general-service/type/request"
	"github.com/dailoi280702/vrs-general-service/usecase/video"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite

	ctx           context.Context
	uc            *video.Usecase
	mockVideoRepo *mockvideo.MockI
}

func (s *TestSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockVideoRepo = &mockvideo.MockI{}

	s.uc = &video.Usecase{
		VideoRepo: s.mockVideoRepo,
	}
}

func TestVideoUsecase(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) TestGetVideoById_Success() {
	req := request.GetById{
		Id: 1,
	}
	expectedVideo := model.Video{
		ID:        1,
		Name:      "Test Video",
		Likes:     100,
		Comments:  50,
		Shares:    20,
		Length:    3600,
		WatchTime: 1800,
	}

	s.mockVideoRepo.On("Get", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(expectedVideo, nil)

	video, err := s.uc.GetVideoById(s.ctx, req)

	s.NoError(err)
	s.Equal(expectedVideo, video)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestGetVideoById_ValidationError() {
	req := request.GetById{
		Id: 0, // Invalid ID
	}

	_, err := s.uc.GetVideoById(s.ctx, req)

	s.Error(err)
}

func (s *TestSuite) TestGetVideoById_NotFoundError() {
	req := request.GetById{
		Id: 1,
	}

	s.mockVideoRepo.On("Get", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(model.Video{}, gorm.ErrRecordNotFound)

	_, err := s.uc.GetVideoById(s.ctx, req)

	s.Error(err)
	s.Equal(video.ErrVideoNotFound, err)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestGetVideoById_RepositoryError() {
	req := request.GetById{
		Id: 1,
	}

	s.mockVideoRepo.On("Get", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(model.Video{}, errors.New("repository error"))

	_, err := s.uc.GetVideoById(s.ctx, req)

	s.Error(err)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestGetVideoByIds_Success() {
	req := request.GetByIds{
		Ids: []int64{1, 2, 3},
	}
	expectedVideos := []model.Video{
		{ID: 1, Name: "Video 1"},
		{ID: 2, Name: "Video 2"},
		{ID: 3, Name: "Video 3"},
	}

	s.mockVideoRepo.On("Find", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(expectedVideos, nil)

	videos, err := s.uc.GetVideoByIds(s.ctx, req)

	s.NoError(err)
	s.Equal(expectedVideos, videos)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestGetVideoByIds_RepositoryError() {
	req := request.GetByIds{
		Ids: []int64{1, 2, 3},
	}

	s.mockVideoRepo.On("Find", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(nil, errors.New("repository error"))

	_, err := s.uc.GetVideoByIds(s.ctx, req)

	s.Error(err)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestGetVideoByUserHistory_Success() {
	req := request.GetById{
		Id: 1,
	}
	expectedVideos := []model.Video{
		{ID: 1, Name: "Video 1"},
		{ID: 2, Name: "Video 2"},
	}

	s.mockVideoRepo.On("Find", s.ctx, mock.AnythingOfType("*specification.videoByUserHisotry")).Return(expectedVideos, nil)

	videos, err := s.uc.GetVideoByUserHistory(s.ctx, req)

	s.NoError(err)
	s.Equal(expectedVideos, videos)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestGetVideoByUserHistory_ValidationError() {
	req := request.GetById{
		Id: 0, // Invalid ID
	}

	_, err := s.uc.GetVideoByUserHistory(s.ctx, req)

	s.Error(err)
}

func (s *TestSuite) TestGetVideoByUserHistory_RepositoryError() {
	req := request.GetById{
		Id: 1,
	}

	s.mockVideoRepo.On("Find", s.ctx, mock.AnythingOfType("*specification.videoByUserHisotry")).Return(nil, errors.New("repository error"))

	_, err := s.uc.GetVideoByUserHistory(s.ctx, req)

	s.Error(err)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestUpdate_Success() {
	req := request.UpdateVideo{
		Id:        1,
		Name:      "Updated Video Name",
		Likes:     200,
		Comments:  100,
		Shares:    50,
		Length:    7200,
		WatchTime: 3600,
	}

	existingVideo := model.Video{
		ID:        1,
		Name:      "Original Video Name",
		Likes:     100,
		Comments:  50,
		Shares:    25,
		Length:    3600,
		WatchTime: 1800,
	}

	s.mockVideoRepo.On("Get", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(existingVideo, nil)
	s.mockVideoRepo.On("Update", s.ctx, mock.AnythingOfType("model.Video")).Return(nil)

	err := s.uc.Update(s.ctx, req)

	s.NoError(err)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestUpdate_ValidationError() {
	req := request.UpdateVideo{
		Id: 0, // Invalid ID
	}

	err := s.uc.Update(s.ctx, req)

	s.Error(err)
}

func (s *TestSuite) TestUpdate_NotFoundError() {
	req := request.UpdateVideo{
		Id:        1,
		Name:      "Updated Video Name",
		Likes:     200,
		Comments:  100,
		Shares:    50,
		Length:    7200,
		WatchTime: 3600,
	}

	s.mockVideoRepo.On("Get", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(model.Video{}, gorm.ErrRecordNotFound)

	err := s.uc.Update(s.ctx, req)

	s.Error(err)
	s.Equal(video.ErrVideoNotFound, err)
	s.mockVideoRepo.AssertExpectations(s.T())
}

func (s *TestSuite) TestUpdate_RepositoryError() {
	req := request.UpdateVideo{
		Id:        1,
		Name:      "Updated Video Name",
		Likes:     200,
		Comments:  100,
		Shares:    50,
		Length:    7200,
		WatchTime: 3600,
	}

	existingVideo := model.Video{
		ID:        1,
		Name:      "Original Video Name",
		Likes:     100,
		Comments:  50,
		Shares:    25,
		Length:    3600,
		WatchTime: 1800,
	}

	s.mockVideoRepo.On("Get", s.ctx, mock.AnythingOfType("*specification.videoByIds")).Return(existingVideo, nil)
	s.mockVideoRepo.On("Update", s.ctx, mock.AnythingOfType("model.Video")).Return(errors.New("repository error"))

	err := s.uc.Update(s.ctx, req)

	s.Error(err)
	s.mockVideoRepo.AssertExpectations(s.T())
}
