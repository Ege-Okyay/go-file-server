package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AllowedOrigin string
	UploadPath    string
	UploadURL     string
	Port          string
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	AllowedOrigin = os.Getenv("ALLOWED_ORIGIN")
	UploadPath = os.Getenv("UPLOAD_PATH")
	UploadURL = os.Getenv("UPLOAD_URL")
	Port = os.Getenv("PORT")
}
