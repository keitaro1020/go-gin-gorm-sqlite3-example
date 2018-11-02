package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/domain"
)

type BookRepository interface {
	Create(db *gorm.DB, b *domain.Book) (*domain.Book, error)
	FindById(db *gorm.DB, id uint) (*domain.Book, error)
	FindAll(db *gorm.DB) ([]*domain.Book, error)
	Update(db *gorm.DB, b *domain.Book) (*domain.Book, error)
	Delete(db *gorm.DB, id uint) error
}

type BookRepositoryImpl struct{}

func GetBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

type BookRecord struct {
	gorm.Model
	Title  string
	Author string
	Price  uint32
}

func (r *BookRecord) TableName() string {
	return "books"
}

func (re *BookRepositoryImpl) Create(db *gorm.DB, b *domain.Book) (*domain.Book, error) {
	r := re.toRecord(b)

	if cr := db.Create(&r); cr.Error != nil {
		return nil, cr.Error
	}

	return re.fromRecord(r), nil
}

func (re *BookRepositoryImpl) FindById(db *gorm.DB, id uint) (*domain.Book, error) {
	r := &BookRecord{}

	if re := db.First(r, id); re.Error != nil {
		return nil, re.Error
	}

	return re.fromRecord(r), nil
}

func (re *BookRepositoryImpl) FindAll(db *gorm.DB) ([]*domain.Book, error) {
	var rs []BookRecord
	if re := db.Find(&rs); re.Error != nil {
		return nil, re.Error
	}

	var bs []*domain.Book
	for _, r := range rs {
		bs = append(bs, re.fromRecord(&r))
	}

	return bs, nil
}

func (re *BookRepositoryImpl) Update(db *gorm.DB, b *domain.Book) (*domain.Book, error) {
	br := re.toRecord(b)

	if re := db.Model(br).Update(br); re.Error != nil {
		return nil, re.Error
	}

	return re.fromRecord(br), nil
}

func (re *BookRepositoryImpl) Delete(db *gorm.DB, id uint) error {
	if re := db.Delete(&BookRecord{Model: gorm.Model{ID: id}}); re.Error != nil {
		return re.Error
	}
	return nil
}

func (re *BookRepositoryImpl) toRecord(b *domain.Book) *BookRecord {
	return &BookRecord{
		Model: gorm.Model{
			ID: b.ID,
		},
		Title:  b.Title,
		Author: b.Author,
		Price:  b.Price,
	}
}

func (re *BookRepositoryImpl) fromRecord(b *BookRecord) *domain.Book {
	return &domain.Book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Price:  b.Price,
	}
}
