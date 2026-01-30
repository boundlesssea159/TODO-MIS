package application

import (
	"TODO-MIS/adapter/driving/api/dto"
	"TODO-MIS/domain/todo"
	"TODO-MIS/domain/todo/entity"
	"TODO-MIS/domain/todo/mock"
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
	service *todo.Todo
	logger  *zap.Logger
}

func (s *TodoTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock.NewMockTodoRepository(s.ctrl)
	s.logger = zap.NewNop()
	s.service = todo.NewTodo(s.repo, s.logger)
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
	s.repo.EXPECT().Create(gomock.Any(), req.Title, req.Description, gomock.Any()).Return(1, nil)
	id, err := s.app.Create(s.ctx, req, 100023323)
	s.Nil(err)
	s.Equal(1, id)
}

func (s *TodoTestSuite) TestCreate_Fail() {
	req := &dto.CreateTodoRequest{
		Title:       "test",
		Description: "desc",
	}
	s.repo.EXPECT().Create(gomock.Any(), req.Title, req.Description, gomock.Any()).Return(0, errors.New("duplicate"))
	id, err := s.app.Create(s.ctx, req, 100023323)
	s.NotNil(err)
	s.Equal(0, id)
}

func (s *TodoTestSuite) TestDelete_Success() {
	id := 1
	s.repo.EXPECT().Delete(gomock.Any(), id, 11212121).Return(nil)
	err := s.app.Delete(s.ctx, id, 11212121)
	s.Nil(err)
}

func (s *TodoTestSuite) TestDelete_Fail() {
	id := 1
	s.repo.EXPECT().Delete(gomock.Any(), id, 11212121).Return(errors.New("not found"))
	err := s.app.Delete(s.ctx, id, 11212121)
	s.NotNil(err)
	s.Equal("not found", err.Error())
}

func (s *TodoTestSuite) TestList_Success() {
	expectedItems := []*entity.TodoItem{
		{
			ID:          1,
			Title:       "Test Todo",
			Description: "Test Description",
			Status:      0,
		},
		{
			ID:          2,
			Title:       "Another Todo",
			Description: "Another Description",
			Status:      1,
		},
	}
	s.repo.EXPECT().List(gomock.Any(), gomock.Any()).Return(expectedItems, nil)
	items, err := s.app.List(s.ctx, 100023323)
	s.Nil(err)
	s.Equal(len(expectedItems), len(items))
	s.Equal(expectedItems[0].ID, items[0].ID)
	s.Equal(expectedItems[0].Title, items[0].Title)
	s.Equal(expectedItems[1].ID, items[1].ID)
	s.Equal(expectedItems[1].Title, items[1].Title)
}

func (s *TodoTestSuite) TestList_Fail() {
	s.repo.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, errors.New("database error"))
	items, err := s.app.List(s.ctx, 0)
	s.NotNil(err)
	s.Equal("database error", err.Error())
	s.Nil(items)
}

func (s *TodoTestSuite) TestComplete_Success() {
	id := 1
	s.repo.EXPECT().Complete(gomock.Any(), id, 11212121).Return(nil)
	err := s.app.Complete(s.ctx, id, 11212121)
	s.Nil(err)
}

func (s *TodoTestSuite) TestComplete_Fail() {
	id := 1
	s.repo.EXPECT().Complete(gomock.Any(), id, 11212121).Return(errors.New("not found"))
	err := s.app.Complete(s.ctx, id, 11212121)
	s.NotNil(err)
	s.Equal("not found", err.Error())
}

func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoTestSuite))
}
