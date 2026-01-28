package persistence

import (
	_const "TODO-MIS/common/const"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MysqlRepositoryTestSuite struct {
	suite.Suite
	db     *gorm.DB
	repo   *MysqlRepository
	logger *zap.Logger
	ctx    context.Context
}

func (s *MysqlRepositoryTestSuite) SetupTest() {
	s.ctx = context.Background()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.Require().NoError(err)

	err = db.AutoMigrate(&TodoItem{})
	s.Require().NoError(err)

	s.db = db
	s.logger, _ = zap.NewDevelopment()
	s.repo = &MysqlRepository{
		db:     s.db,
		logger: s.logger,
	}
}

func (s *MysqlRepositoryTestSuite) TestCreate_Success() {
	title := "Test Title"
	description := "Test Description"
	userId := 100023323

	id, err := s.repo.Create(s.ctx, title, description, userId)

	s.NoError(err)
	s.Greater(id, 0)

	var item TodoItem
	result := s.db.First(&item, id)
	s.NoError(result.Error)
	s.Equal(title, item.Title)
	s.Equal(description, item.Description)
	s.Equal(_const.TodoItemCreatedStatus, item.Status)
	s.Equal(userId, item.UserID)
}

func (s *MysqlRepositoryTestSuite) TestCreate_Fail() {
	sqlDB, err := s.db.DB()
	s.Require().NoError(err)
	s.Require().NoError(sqlDB.Close())

	_, err = s.repo.Create(s.ctx, "title", "desc", 1)

	s.Error(err)
}

func (s *MysqlRepositoryTestSuite) TestDelete_Success() {
	item := &TodoItem{
		Title:       "To Delete",
		Description: "Description to delete",
		Status:      _const.TodoItemCreatedStatus,
		UserID:      100023323,
	}
	s.db.Create(item)

	err := s.repo.Delete(s.ctx, item.ID)
	s.NoError(err)

	var updatedItem TodoItem
	result := s.db.First(&updatedItem, item.ID)
	s.NoError(result.Error)
	s.Equal(_const.TodoItemDeletedStatus, updatedItem.Status)
}

func (s *MysqlRepositoryTestSuite) TestDelete_Fail() {
	item := &TodoItem{
		Title:       "Another Item",
		Description: "Another Description",
		Status:      _const.TodoItemCreatedStatus,
		UserID:      100023323,
	}
	s.db.Create(item)

	err := s.repo.Delete(s.ctx, item.ID)
	s.NoError(err)

	err = s.repo.Delete(s.ctx, 999999)
	s.NoError(err)
}

func (s *MysqlRepositoryTestSuite) TestList_Success() {
	userID := 100023323
	items := []TodoItem{
		{
			Title:       "First Item",
			Description: "First Description",
			Status:      _const.TodoItemCreatedStatus,
			UserID:      userID,
		},
		{
			Title:       "Second Item",
			Description: "Second Description",
			Status:      _const.TodoItemCreatedStatus,
			UserID:      userID,
		},
	}
	for _, item := range items {
		s.db.Create(&item)
	}

	results, err := s.repo.List(s.ctx, userID)

	s.NoError(err)
	s.Len(results, 2)

	s.Contains([]string{results[0].Title, results[1].Title}, "First Item")
	s.Contains([]string{results[0].Title, results[1].Title}, "Second Item")
}

func (s *MysqlRepositoryTestSuite) TestList_Empty() {
	results, err := s.repo.List(s.ctx, 999999)
	s.NoError(err)
	s.Empty(results)
}

func (s *MysqlRepositoryTestSuite) TestList_Fail() {
	// 关闭数据库连接以模拟错误
	sqlDB, err := s.db.DB()
	s.Require().NoError(err)
	s.Require().NoError(sqlDB.Close())
	_, err = s.repo.List(s.ctx, 1)
	s.Error(err)
}

func (s *MysqlRepositoryTestSuite) TestComplete_Success() {
	item := &TodoItem{
		Title:       "To Complete",
		Description: "Description to complete",
		Status:      _const.TodoItemCreatedStatus,
		UserID:      100023323,
	}
	s.db.Create(item)
	err := s.repo.Complete(s.ctx, item.ID)
	s.NoError(err)
	var updatedItem TodoItem
	result := s.db.First(&updatedItem, item.ID)
	s.NoError(result.Error)
	s.Equal(_const.TodoItemDoneStatus, updatedItem.Status)
}

func (s *MysqlRepositoryTestSuite) TestComplete_Fail() {
	err := s.repo.Complete(s.ctx, 999999)
	s.NoError(err)
}

func (s *MysqlRepositoryTestSuite) TearDownTest() {
	if s.logger != nil {
		s.logger.Sync()
	}
}

func TestMysqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MysqlRepositoryTestSuite))
}
