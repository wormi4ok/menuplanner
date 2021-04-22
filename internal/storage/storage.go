package storage

import (
	"github.com/wormi4ok/menuplanner/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func InitDB(dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&internal.Recipe{}, &Week{}); err != nil {
		return nil, err
	}

	instance := &DB{db: db}

	instance.preloadCourses()

	return instance, nil
}
