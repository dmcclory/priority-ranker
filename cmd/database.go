package cmd

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
	"fmt"
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
		return &gorm.DB{}, fmt.Errorf("loadDb: failed executing query: %w", err)
	}
	return db, nil
}

func initDb(path string) (*gorm.DB, error) {
	db, err := loadDb(path)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("initDb: failed executing query: %w", err)
	}

	db.AutoMigrate(&Option{})
	db.AutoMigrate(&Vote{})

	return db, nil
}

func loadOptions(db *gorm.DB) ([]Option, error) {
	var options []Option
	err := db.Find(&options).Error
	if err != nil {
		return []Option{}, fmt.Errorf("loadOptions: failed executing query: %w", err)
	}
	return options, nil
}

func getOption(db *gorm.DB, optionId uint) (Option, error) {
	var option Option

	err := db.First(&option, optionId).Error
	if err != nil {
		return Option{}, fmt.Errorf("getOption: failed executing query: %w", err)
	}
	return option, nil
}

func addOption(db *gorm.DB, newOption string) (Option, error) {
	var option Option
	err := db.FirstOrCreate(&option, Option{Label: newOption}).Error
	if err != nil {
		return Option{}, fmt.Errorf("addOption: failed executing query: %w", err)
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
	  return[]Option{}, fmt.Errorf("addOptions: failed executing query: %w", err)
	}
	return optionInputs, nil
}

func removeOption(db *gorm.DB, optionId uint) error {
	// https://gorm.io/docs/error_handling.html
	// "After a chain of methods, it’s crucial to check the Error field"
	err := db.Delete(&Option{}, optionId).Error

	if err != nil {
		return fmt.Errorf("removeOption: failed executing query: %w", err)
	}
	return nil
}

func addVote(db *gorm.DB, winnerId uint, loserId uint) error {
	createdAt := time.Now().Unix()
	vote := Vote{WinnerId: winnerId, LoserId: loserId, CreatedAt: createdAt}

	err := db.Create(&vote).Error

	if err != nil {
		return fmt.Errorf("addVote: failed executing query: %w", err)
	}
	return nil
}

func loadVotes(db *gorm.DB) ([]Vote, error) {
	var votes []Vote
	err := db.Find(&votes).Error

	if err != nil {
	  return[]Vote{}, fmt.Errorf("loadVotes: failed executing query: %w", err)
	}
	return votes, nil
}

func deleteVotes(db *gorm.DB, optionId uint) (int64, error) {
	result := db.Where("winner_id = ? or loser_id = ?", optionId, optionId).Delete(&Vote{})

	err := result.Error

	if err != nil {
	  return 0, fmt.Errorf("deleteVotes: failed executing query: %w", err)
	}

	return result.RowsAffected, nil
}

func deleteOptionAndVotes(db *gorm.DB, optionId uint) error {
	tx := db.Begin()

	_, err := deleteVotes(tx, optionId)

	if err != nil {
		tx.Rollback()
	  return fmt.Errorf("deleteOptionAndVotes: failed executing query: %w", err)
	}

	err = removeOption(tx, optionId)

	if err != nil {
		tx.Rollback()
	  return fmt.Errorf("deleteOptionAndVotes: failed executing query: %w", err)
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
	  return fmt.Errorf("deleteOptionAndVotes: failed executing query: %w", err)
	}

	return nil
}
