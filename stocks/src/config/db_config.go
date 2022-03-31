package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConfig(dbHost, dbPort, dbName, dbUser, dbPassword string) {
	datasourceUrl := GetDbUrl(dbHost, dbPort, dbName, dbUser, dbPassword)

	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  datasourceUrl,
			PreferSimpleProtocol: true,
		},
	), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Db = db

}

func GetDbUrl(dbHost, dbPort, dbName, dbUser, dbPassword string) string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPassword,
	)
}
