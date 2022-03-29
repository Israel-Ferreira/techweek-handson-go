package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")

	DbName = os.Getenv("DB_NAME")
	DbUser = os.Getenv("DB_USER")

	DbPassword = os.Getenv("DB_PASSWORD")

}

func InitConfig() {
	LoadEnv()
	DbConfig(DbHost, DbPort, DbName, DbUser, DbPassword)
}
