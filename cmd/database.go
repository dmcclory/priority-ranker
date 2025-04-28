package cmd

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
	"gorm.io/gorm/logger"
)

type Option struct {
	ID    uint
	Label string
}

type Vote struct {
	WinnerId uint
	LoserId uint
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
	db.AutoMigrate(&Vote{})

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

func addVote(db *gorm.DB, winnerId uint, loserId uint) error {
	createdAt := time.Now().Unix()
	vote := Vote{WinnerId: winnerId, LoserId: loserId, CreatedAt: createdAt}

	result := db.Create(&vote).Error

	return result
}

func loadVotes(db *gorm.DB) ([]Vote, error) {
	var votes []Vote
	err := db.Find(&votes).Error

	return votes, err
}
