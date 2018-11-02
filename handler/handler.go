package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	Config *Config
}

type Config struct {
}

type ErrorResponse struct {
	Message    string      `json:"message"`
	InnerCode  int         `json:"code,omitempty"`
	StatusCode int         `json:"-"`
	Validates  []*Validate `json:"validates,omitempty"`
}

type Validate struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

var (
	db *gorm.DB
)

func SetDb(d *gorm.DB) {
	db = d
}

func NewHandler(cfg *Config) (*Handler, error) {
	return &Handler{
		Config: cfg,
	}, nil
}

func NewErrorResponse(err error) *ErrorResponse {
	e := &ErrorResponse{
		StatusCode: ErrorCode(err),
	}

	if err != nil {
		e.Message = err.Error()

		switch err.(type) {
		case *json.UnmarshalTypeError:
			v := err.(*json.UnmarshalTypeError)
			e.Message = fmt.Sprintf("Field type of the %s is a %s", v.Field, v.Type)
		case *time.ParseError:
			e.Message = "Field type of the time is parse error"
		}

		if errs, ok := err.(validator.ValidationErrors); ok {
			e.Message = "Validation error"
			e.Validates = []*Validate{}

			for _, v := range errs {
				e.Validates = append(e.Validates, &Validate{
					Key:     v.Field,
					Message: fmt.Sprintf("Field validation for %s failed on the %s", v.Field, v.ActualTag),
				})
			}
		}
	}

	return e
}

func (e *ErrorResponse) render(gc *gin.Context) {
	if e.StatusCode >= http.StatusInternalServerError {
		log.Printf("%s", e.Message)
	}

	gc.JSON(e.StatusCode, e)
}

func ErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err.(type) {
	case *strconv.NumError:
		return http.StatusBadRequest
	case *json.UnmarshalTypeError:
		return http.StatusBadRequest
	case *time.ParseError:
		return http.StatusBadRequest
	case validator.ValidationErrors:
		return http.StatusBadRequest
	}

	switch err.Error() {
	case gorm.ErrRecordNotFound.Error():
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

func (h *Handler) Transaction(c *gin.Context, txFunc func(*gorm.DB, *gin.Context) error) (err error) {
	tx := db.Begin()
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			log.Printf("ROLLBACK! [err : %v]", err)
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	return txFunc(tx, c)
}
