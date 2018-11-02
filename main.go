package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/handler"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/middleware"
	"github.com/keitaro1020/go-gin-gorm-sqlite3-example/repository"
	"gopkg.in/go-playground/validator.v8"
)

type Config struct {
	AppConfig      *AppConfig         `bind:"required"`
	HandlerConfig  *handler.Config    `bind:"required"`
	DatabaseConfig *repository.Config `bind:"required"`
}

type AppConfig struct {
	Mode string `bind:"required"`
	Port int32  `bind:"required"`
}

var cfg *Config

func main() {
	err := initialize()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.DbConnectionMiddleware())

	h, err := handler.NewHandler(cfg.HandlerConfig)
	if err != nil {
		panic(err)
	}

	// hello world
	r.GET("/hello", h.Hello)

	// book
	r.POST("/book", h.CreateBook)
	r.GET("/book", h.GetBooks)

	r.Run(fmt.Sprintf(":%v", cfg.AppConfig.Port))
}

func initialize() error {
	err := readConfig()
	if err != nil {
		return err
	}

	gin.SetMode(cfg.AppConfig.Mode)

	repository.SetConfig(cfg.DatabaseConfig)

	return nil
}

func readConfig() error {
	cfg = &Config{}

	if _, err := toml.DecodeFile("app.conf", &cfg); err != nil {
		return err
	}

	if err := validator.New(&validator.Config{}).Struct(cfg); err != nil {
		return err
	}

	return nil
}
