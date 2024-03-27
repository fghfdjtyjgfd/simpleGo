package connection

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config config
type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Debug        bool
}

type Session struct {
	Db *gorm.DB
}

func New(c *Config) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s  sslmode=disable password=",
		c.Host,
		c.Port,
		c.User,
		c.DatabaseName,
	)

	dns = dns + string(c.Password)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
