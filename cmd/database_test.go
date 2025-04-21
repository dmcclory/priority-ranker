package cmd

import (
	"os"
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
