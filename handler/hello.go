package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *HandlerImpl) Hello(gc *gin.Context) {
	res := &struct {
		Value string
	}{
		Value: "Hello World",
	}
	gc.JSON(http.StatusOK, res)
	log.Print("hello world")
}
