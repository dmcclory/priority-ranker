package cmd

import (
  "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Option struct {
	ID uint
	Label string
}

func loadDb(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	check(err)
	return db
}

func initDb(path string) *gorm.DB {
	db := loadDb(path)

	db.AutoMigrate(&Option{})

	return db
}

func loadOptions(db *gorm.DB) []Option {
	var options []Option
	db.Find(&options)
	return options
}
