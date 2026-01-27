package application

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/domain"
	"TODO-MIS/domain/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type TodoTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	repo    *mock.MockTodoRepository
	app     *Todo
	ctx     context.Context
	service *domain.Todo
	logger  *zap.Logger
}

func (s *TodoTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock.NewMockTodoRepository(s.ctrl)
	s.logger = zap.NewNop()
	s.service = domain.NewTodo(s.repo, s.logger)
	s.app = &Todo{
		service: s.service,
		logger:  s.logger,
	}
}

func (s *TodoTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *TodoTestSuite) TestCreate_Success() {
	req := &dto.CreateTodoRequest{
		Title:       "test",
		Description: "desc",
	}
	s.repo.EXPECT().Create(gomock.Any(), req.Title, req.Description).Return(1, nil)
	id, err := s.app.Create(s.ctx, req)
	s.Nil(err)
	s.Equal(1, id)
}

func (s *TodoTestSuite) TestCreate_Fail() {
	req := &dto.CreateTodoRequest{
		Title:       "test",
		Description: "desc",
	}
	s.repo.EXPECT().Create(gomock.Any(), req.Title, req.Description).Return(0, errors.New("duplicate"))
	id, err := s.app.Create(s.ctx, req)
	s.NotNil(err)
	s.Equal(0, id)
}

func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoTestSuite))
}
