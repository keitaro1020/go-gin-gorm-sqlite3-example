package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/handler"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/repository"
)

func DbConnectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := repository.ConnectDB()

		if err != nil {
			panic(err)
		}

		defer db.Close()

		handler.SetDb(db)

		db.LogMode(true)

		c.Next()
	}
}
