package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "9010"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBAddress: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "mariadb"),
			getEnv("DB_PORT", "3306"),
		),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTSecret:              getEnv("JWT_SECRET", "secret?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return fallback
	}
}

func getEnvAsInt(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)
	if ok {

		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		} else {
			return value
		}

	} else {
		return fallback
	}
}
