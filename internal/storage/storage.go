package storage

import (
	"log"
	"os"
	"time"

	"github.com/wormi4ok/menuplanner/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	db *gorm.DB
}

func InitDB(dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: customLogger(),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&internal.Recipe{}, &Week{}, &internal.User{}); err != nil {
		return nil, err
	}

	instance := &DB{db: db}

	instance.preloadCourses()

	return instance, nil
}

func customLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}
