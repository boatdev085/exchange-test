package connection

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDb(connectionString string) *gorm.DB {
	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDbInstance() *gorm.DB {
	return DB
}
