package config

import (
	"os"

	"github.com/trickaugusto/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("SQLite database not found, creating new one at %s", dbPath)
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Error creating SQLite database: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Error creating SQLite database: %v", err)
			return nil, err
		}

		file.Close()
	}

	// create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating SQLite database: %v", err)
		return nil, err
	}

	return db, nil

}
