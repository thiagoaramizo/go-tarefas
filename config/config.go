package config

import "gorm.io/gorm"

var ( 
	db *gorm.DB
	logger *Logger 
)

func Init() error {
	var err error
	db, err = InitializaSQLite()
	if err != nil {
		return err
	}
	return nil
}

func GetDataBase() *gorm.DB {
	return db
}


func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}