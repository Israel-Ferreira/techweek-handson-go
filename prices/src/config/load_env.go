package config

import (
	"log"
	"os"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/models"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

var (
	DbHost      string
	DbPort      string
	DbName      string
	DbUser      string
	DbPassword  string
	KafkaServer string
	KafkaTopic  string

	KafkaProductConsumer *kafka.Reader
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado")
	}

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")

	DbName = os.Getenv("DB_NAME")
	DbUser = os.Getenv("DB_USER")

	DbPassword = os.Getenv("DB_PASSWORD")
	KafkaServer = os.Getenv("KAFKA_SERVER")
	KafkaTopic = os.Getenv("KAFKA_TOPIC")

}

func InitConfig() {
	LoadEnv()

	DbConfig(DbHost, DbPort, DbName, DbUser, DbPassword)

	KafkaProductConsumer = ConsumerConfig(KafkaServer, KafkaTopic, "price")

	Db.AutoMigrate(&models.Price{})
}
