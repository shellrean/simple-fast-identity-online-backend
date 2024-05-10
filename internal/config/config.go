package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error when load .env: ", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
