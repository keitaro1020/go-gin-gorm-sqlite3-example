package handler

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/domain"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/repository"
)

func TestCreateBook(t *testing.T) {

	var res *httptest.ResponseRecorder
	var err error

	bk := &domain.Book{
		Title: "test title",
		Price: 1234,
		Author: "test author",
	}

	bMock := new(repository.BookRepositoryMock)
	bMock.On("Create", db, bk).Return(bk, nil)

	testHandler := HandlerImpl{
		config:         &Config{},
		bookRepository: bMock,
	}

	r := &CreateBookRequest{
		Title: bk.Title,
		Price: bk.Price,
		Author: bk.Author,
	}
	res, err = Do(&Pattern{
		Request:        NewRequest(http.MethodPost, JsonEncode(r)),
		HandlerFunc:    testHandler.CreateBook,
		WantStatusCode: http.StatusOK,
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	b := &domain.Book{}
	json.Unmarshal(res.Body.Bytes(), b)

	assert.Equal(t, b, bk)
}
