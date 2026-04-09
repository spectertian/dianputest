package config

import "os"

var (
	DSN       = getEnv("DSN", "root:123456@tcp(127.0.0.1:3306)/dianpu?charset=utf8mb4&parseTime=True&loc=Local")
	JWTSecret = getEnv("JWT_SECRET", "dianpu-secret-key")
	Port      = getEnv("PORT", ":8080")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
