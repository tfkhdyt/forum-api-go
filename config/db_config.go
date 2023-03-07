package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ====================

type dbConfig struct {
	host     string
	username string
	password string
	name     string
	port     string
}

// ====================

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}
}

// ====================

func GetPostgresConnectionString() string {
	postgresDBConfig := dbConfig{
		host:     os.Getenv("DB_HOST"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		name:     os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		postgresDBConfig.host,
		postgresDBConfig.username,
		postgresDBConfig.password,
		postgresDBConfig.name,
		postgresDBConfig.port,
	)

	return dsn
}
