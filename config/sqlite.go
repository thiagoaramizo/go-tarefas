package config

import (
	"os"

	"github.com/thiagoaramizo/go-tarefas/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializaSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("Database file does not exist. Creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.ErrorF("Error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.ErrorF("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error initializing SQLite: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schema.Tarefa{})
	if err != nil {
		logger.ErrorF("Error migrating schema: %v", err)
		return nil, err
	}

	return db, nil
}