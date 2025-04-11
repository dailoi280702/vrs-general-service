package auth_test

import (
	"context"
	"errors"
	"testing"

	mockuser "github.com/dailoi280702/vrs-general-service/mock/user"
	"github.com/dailoi280702/vrs-general-service/type/request"
	"github.com/dailoi280702/vrs-general-service/type/response"
	"github.com/dailoi280702/vrs-general-service/usecase/auth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	ctx          context.Context
	uc           *auth.Usecase
	mockUserRepo *mockuser.MockI
}

func (s *TestSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockUserRepo = &mockuser.MockI{}

	s.uc = &auth.Usecase{
		UserRepo: s.mockUserRepo,
	}
}

func TestAuth(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) TestRegister_Success() {
	req := request.Register{
		Email:    "test@test.com",
		FullName: "test",
		Password: "12345678",
	}

	expectedResp := response.Register{
		Email:    "test@test.com",
		FullName: "test",
	}

	ctx := context.Background()

	s.mockUserRepo.On("Count", ctx, mock.Anything).Return(int64(0), nil)
	s.mockUserRepo.On("Create", ctx, mock.Anything).Return(nil)

	resp, err := s.uc.Register(ctx, req)

	s.Nil(err)
	s.Equal(expectedResp, resp)
}

func (s *TestSuite) TestRegister_Validation() {
	req := request.Register{
		Email:    "",
		FullName: "test",
		Password: "12345678",
	}

	ctx := context.Background()

	_, err := s.uc.Register(ctx, req)

	s.NotNil(err)
}

func (s *TestSuite) TestRegister_CountUser() {
	req := request.Register{
		Email:    "test@test.com",
		FullName: "test",
		Password: "12345678",
	}

	ctx := context.Background()

	s.mockUserRepo.On("Count", ctx, mock.Anything).Return(int64(0), errors.New("test"))

	_, err := s.uc.Register(ctx, req)

	s.NotNil(err)
}

func (s *TestSuite) TestRegister_CreateUser() {
	req := request.Register{
		Email:    "test@test.com",
		FullName: "test",
		Password: "12345678",
	}

	ctx := context.Background()

	s.mockUserRepo.On("Count", ctx, mock.Anything).Return(int64(0), nil)
	s.mockUserRepo.On("Create", ctx, mock.Anything).Return(errors.New("test"))

	_, err := s.uc.Register(ctx, req)

	s.NotNil(err)
}
