package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string 
	JWTSecret string 
}


func LoadConfig() *Config {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	
	return &Config{
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}