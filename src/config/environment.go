package config

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTPrivateKey string
	JWTPublicKey  string
}

var ENV EnvConfig

func InitEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Decode jwt
	JWTPrivateKey, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_PRIVATE_KEY_BASE64_ENCODED"))
	if err != nil {
		panic(DECODE_JWT_PRIVATE_KEY_ERROR)
	}
	JWTPublicKey, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_PUBLIC_KEY_BASE64_ENCODED"))
	if err != nil {
		panic(DECODE_JWT_PUBLIC_KEY_ERROR)
	}

	ENV = EnvConfig{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		JWTPrivateKey: string(JWTPrivateKey),
		JWTPublicKey:  string(JWTPublicKey),
	}

}
