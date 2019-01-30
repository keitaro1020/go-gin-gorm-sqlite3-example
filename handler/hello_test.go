package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestHello(t *testing.T) {

	var res *httptest.ResponseRecorder
	var err error

	testHandler := HandlerImpl{
		config:         &Config{},
	}

	res, err = Do(&Pattern{
		Request:        NewRequest(http.MethodGet, ""),
		HandlerFunc:    testHandler.Hello,
		WantStatusCode: http.StatusOK,
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	b := &struct {
		Value string
	}{
	}
	json.Unmarshal(res.Body.Bytes(), b)

	assert.Equal(t, b.Value, "Hello World")
}
