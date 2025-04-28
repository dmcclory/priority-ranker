package cmd

import (
	"os"
	"time"
	"testing"
)

func TestInitDbCreatesASqliteDbIfItDoesNotExist(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	_, err = initDb(dbPath("test-db"))
	check(err)

	if fileDoesNotExist(dbPath("test-db")) {
		t.Errorf("expected to find test-db database in the working directory")
	}
}

func TestInitDbCreatesAnOptionsTableInNewDb(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	db, err := initDb(dbPath("test-db"))
	check(err)

	var tableNames []string
	included := false

	db.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tableNames)

	if len(tableNames) == 0 {
		t.Errorf("the query to return more than 0 results")
	}

	for _, name := range tableNames {
		if name == "options" {
			included = true
		}
	}

	if !included {
		t.Errorf("expected initDb to create an options table, but it did not")
	}
}

func TestAddNewOptionToDbAddsStringAsANewOption(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	db, err := initDb(dbPath("test-db"))

	addOption(db, "cool test")

	var result Option
	db.Where("label = ?", "cool test").First(&result)

	if result.ID == 0 {
		t.Errorf("no record found in DB, but one was expected")
	}
}

func TestAddNewOptionToDbDoesNotDuplicateRowsIfEntryAlreadyExists(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	db, err := initDb(dbPath("test-db"))

	addOption(db, "cool test")
	addOption(db, "cool test")

	var result Option
	db.Where("label = ?", "cool test").First(&result)

	if result.ID != 1 {
		t.Errorf("no record found in DB, but one was expected")
	}

	var count int64
	db.Table("options").Count(&count)

	if count != 1 {
		t.Error("expected only one row to exist")
	}
}

func TestRemoveOptionFromDbDeletesARow(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	db, err := initDb(dbPath("test-db"))

	addOption(db, "cool test")

	var count int64

	db.Table("options").Count(&count)
	if count != 1 {
		t.Error("expected the table to have one row")
	}

	removeOption(db, 1)


	db.Table("options").Count(&count)
	if count != 0 {
		t.Error("expected the table to be empty")
	}

}

func TestAddVoteSavesAWinnerAndLoserId(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	db, err := initDb(dbPath("test-db"))

	// this works because the granularity of Unix() is 1 second
	// and this will done well within one second
	expectedCreatedAt := time.Now().Unix()
	err = addVote(db, 1, 2)

	if err != nil {
		t.Fatalf("addVote failed with error: %e", err)
	}

	votes, err := loadVotes(db)

	if err != nil {
		t.Fatalf("addVote - failed to load Votes error: %e", err)
	}

	if len(votes) == 0 {
		t.Fatalf("addVote - failed, loadVotes returned an empty list")
	}

	if votes[0].WinnerId != 1 {
		t.Fatalf("addVote: expected WinnerId to equal %d, but got %d", 1, votes[0].WinnerId)
	}

	if votes[0].LoserId != 2 {
		t.Fatalf("addVote: expected LoserId to equal %d, but got %d", 1, votes[0].LoserId)
	}

	if !(expectedCreatedAt == votes[0].CreatedAt) {
		t.Fatalf("addVote: should have a CreatedAt time stamp after %d, but it is %d", expectedCreatedAt, votes[0].CreatedAt)
	}
}

func TestLoadVotesReturnsListofVotes(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	if fileExists(dbPath("test-db")) {
		t.Errorf("test setup failed -found database that should not exist")
	}

	db, err := initDb(dbPath("test-db"))

	votes, err := loadVotes(db)

	if err != nil {
		t.Fatalf("loadVotes failed with error: %e", err)
	}

	if len(votes) != 0 {
		t.Fatalf("loadVotes - expected 0 results, but got %d", len(votes))
	}

	err = addVote(db, 1, 2)
	err = addVote(db, 2, 3)
	err = addVote(db, 3, 4)
	err = addVote(db, 4, 5)

	if err != nil {
		t.Fatalf("addVote failed with error: %e", err)
	}

	votes, err = loadVotes(db)

	if err != nil {
		t.Fatalf("loadVotes failed with error: %e", err)
	}

	if len(votes) != 4 {
		t.Fatalf("loadVotes - expected 4 results, but got %d", len(votes))
	}
}
