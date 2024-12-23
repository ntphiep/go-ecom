package config

import (
	"fmt"
	"os"

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

func initConfig() Config{
  godotenv.Load()
  
  return Config{
    PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
    Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 * 7),
  }
}

func getEnv(key, fallback string) string {
  if value, ok := os.LookUpEnvs(key); ok {
    return value
  }

  return fallback
}
