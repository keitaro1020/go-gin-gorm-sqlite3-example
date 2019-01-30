package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/domain"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) Create(db *gorm.DB, dc *domain.Book) (*domain.Book, error) {
	args := m.Called(db, dc)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) FindById(db *gorm.DB, id uint) (*domain.Book, error) {
	args := m.Called(db, id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) FindAll(db *gorm.DB) ([]*domain.Book, error) {
	args := m.Called(db)
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) Update(db *gorm.DB, b *domain.Book) (*domain.Book, error) {
	args := m.Called(db, b)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) Delete(db *gorm.DB, id uint) error {
	args := m.Called(db, id)
	return args.Error(1)
}
