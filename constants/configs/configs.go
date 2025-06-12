package configs

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Configs struct {
	AppEnv  		string
	AppPort 		string
	AppTokenKey string
	AppHost 		string
	AppURL 			string
	DBHost 			string
	DBPort 			string
	DBUser 			string
	DBPass 			string
	DBName 			string
	DBURL	 			string
	MailUser		string
	MailPass		string
	MailPort		string
	MailHost		string
}

var cfg *Configs
var once sync.Once

func LoadConfig() *Configs {
	once.Do(func () {
		err := godotenv.Load()	
		if err != nil {
			log.Println("Warning: No .env file found, using system environment variables")
		}

		cfg = &Configs{
			AppEnv: getEnv("APP_ENV", "develop"),
			AppPort: getEnv("APP_PORT", "8080"),
			AppTokenKey: getEnv("APP_TOKEN_KEY", ""),
			AppHost: getEnv("APP_HOST", "localhost"),
			AppURL: getEnv("APP_URL", "http://localhost:8080"),
			DBHost: getEnv("DB_HOST", "db"),
			DBPort: getEnv("DB_PORT", "5432"),
			DBUser: getEnv("DB_USER", "username"),
			DBPass: getEnv("DB_PASS", "password"),
			DBName: getEnv("DB_NAME", "user_management"),
			MailUser: getEnv("MAIL_USER", "user@mail.com"),
			MailPass: getEnv("MAIL_PASS", "password"),
			MailPort: getEnv("MAIL_PORT", "587"),
			MailHost: getEnv("MAIL_HOST", "smtp.gmail.com"),
		}

		cfg.DBURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	})

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}