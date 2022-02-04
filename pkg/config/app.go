package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=abul.db.elephantsql.com user=vuzitaei password=0Yflq8UdUrVJhwCVIVuh8TzwKk_Z0sZX dbname=vuzitaei port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetBD() *gorm.DB {
	return db
}
