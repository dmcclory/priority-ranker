package cmd

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Option struct {
	ID    uint
	Label string
}

type Vote struct {
	WinnerId int64
	LoserId int64
	CreatedAt int64
}

func loadDb(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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

func getOption(db *gorm.DB, optionId uint64) (Option, error) {
	var option Option

	err := db.First(&option, optionId).Error

	if err != nil {
		return Option{}, err
	}

	return option, nil
}

func addOption(db *gorm.DB, newOption string) (Option, error) {
	var option Option
	err := db.FirstOrCreate(&option, Option{Label: newOption}).Error
	if err != nil {
		return Option{}, err
	}
	return option, nil
}

func addOptions(db *gorm.DB, newOptions []string) ([]Option, error) {
	var optionInputs []Option

	for _, newOption := range newOptions {
		optionInputs = append(optionInputs, Option{Label: newOption})
	}

	err := db.Create(&optionInputs).Error
	if err != nil {
	  return[]Option{}, err
	}

	return optionInputs, nil
}

func removeOption(db *gorm.DB, optionId uint64) error {
	// https://gorm.io/docs/error_handling.html
	// "After a chain of methods, it’s crucial to check the Error field"
	err := db.Delete(&Option{}, optionId).Error

	return err
}
