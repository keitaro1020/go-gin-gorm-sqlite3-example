package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Config struct {
	DatabaseDialect string `bind:"required"`
	DatabaseUrl     string `bind:"required"`
	MaxIdleConns    int    `bind:"required"`
	MaxOpenConns    int    `bind:"required"`
	ConnMaxLifetime int    `bind:"required"`
}

var (
	config *Config
)

func SetConfig(c *Config) {
	config = c
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(config.DatabaseDialect, config.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

	migrate(db)

	return db, err
}

func migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.AutoMigrate(&BookRecord{})
}
