package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/domain"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/repository"
	"net/http"
)

type CreateBookRequest struct {
	Title  string `json:"title,omitempty" binding:"required"`
	Author string `json:"author,omitempty" binding:"required"`
	Price  uint32 `json:"price,omitempty" binding:"required"`
}

func (h *Handler) CreateBook(gc *gin.Context) {

	req := &CreateBookRequest{}
	if err := gc.Bind(req); err != nil {
		NewErrorResponse(err).render(gc)
		return
	}

	if err := h.Transaction(gc, func(tx *gorm.DB, gc *gin.Context) error {

		br := repository.GetBookRepository()
		bk, err := br.Create(tx, &domain.Book{
			Title:  req.Title,
			Author: req.Author,
			Price:  req.Price,
		})
		if err != nil {
			return err
		}

		gc.JSON(http.StatusCreated, bk)
		return nil
	}); err != nil {
		NewErrorResponse(err).render(gc)
		return
	}
}

func (h *Handler) GetBooks(gc *gin.Context) {
	br := repository.GetBookRepository()
	bks, err := br.FindAll(db)
	if err != nil {
		NewErrorResponse(err).render(gc)
		return
	}

	gc.JSON(http.StatusCreated, bks)
}
