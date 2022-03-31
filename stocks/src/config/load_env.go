package config

import (
	"log"
	"os"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/models"
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

	KafkaConsumerGroup string

	KafkaStockConsumer *kafka.Reader
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

	KafkaServer = os.Getenv("KAFKA_SERVER")
	KafkaTopic = os.Getenv("KAFKA_TOPIC")
	KafkaConsumerGroup = os.Getenv("KAFKA_GROUP_ID")

}

func InitConfig() {
	LoadEnv()
	DbConfig(DbHost, DbPort, DbName, DbUser, DbPassword)

	KafkaStockConsumer = ConsumerConfig(KafkaServer, KafkaTopic, KafkaConsumerGroup)

	Db.AutoMigrate(&models.Stock{})
}
