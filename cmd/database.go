package cmd

import (
  "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Option struct {
	ID uint
	Label string
}

func loadDb(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}
	return db, nil
}

func initDb(path string) (*gorm.DB, error) {
	db, err := loadDb(path)

	if err != nil {
		return &gorm.DB{}, err
	}

	db.AutoMigrate(&Option{})

	return db, nil
}

func loadOptions(db *gorm.DB) []Option {
	var options []Option
	db.Find(&options)
	return options
}
