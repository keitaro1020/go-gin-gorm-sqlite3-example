package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/repository"
)


type Pattern struct {
	Request        *http.Request
	HandlerFunc    gin.HandlerFunc
	WantStatusCode int
}

func TestMain(m *testing.M) {
	var err error

	gin.SetMode(gin.TestMode)

	repository.SetConfig(&repository.Config{
		DatabaseDialect: "sqlite3",
		DatabaseUrl: "testing.sqlite3",
	})
	db, err := repository.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	SetDb(db)

	code := m.Run()

	os.Exit(code)
}

func JsonEncode(v interface{}) string {
	b, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func NewRequest(method, body string) *http.Request {
	r, _ := http.NewRequest(method, "/", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	return r
}

func Do(p *Pattern) (*httptest.ResponseRecorder, error) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)

	r.POST("/", p.HandlerFunc)
	r.ServeHTTP(res, p.Request)

	if res.Code != p.WantStatusCode {
		return nil, errors.New(fmt.Sprintf("Path=%q, StatusCode=%d, Want=%d, ResBody=%s", p.Request.URL.Path, res.Code, p.WantStatusCode, res.Body.String()))
	}

	return res, nil
}
