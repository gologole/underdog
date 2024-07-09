package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env    string
	Server ServerConfig
	DB     DBConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ScheduledShutdown time.Duration
	PeopleInfo        string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type LoggerConfig struct {
	Source string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	env := os.Getenv("ENVIRONMENT")
	serverPort := os.Getenv("SERVER_PORT")
	serverReadTimeout, err := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error parsing SERVER_READ_TIMEOUT: %v", err)
		return nil, err
	}
	serverWriteTimeout, err := time.ParseDuration(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error parsing SERVER_WRITE_TIMEOUT: %v", err)
		return nil, err
	}
	serverScheduledShutdown, err := time.ParseDuration(os.Getenv("SERVER_SCHEDULED_SHUTDOWN"))
	if err != nil {
		log.Fatalf("Error parsing SERVER_SCHEDULED_SHUTDOWN: %v", err)
		return nil, err
	}
	peopleInfo := os.Getenv("PEOPLE_INFO_ADDRESS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	loggerSource := os.Getenv("LOGGER_SOURCE")

	config := &Config{
		Env: env,
		Server: ServerConfig{
			Port:              serverPort,
			ReadTimeout:       serverReadTimeout,
			WriteTimeout:      serverWriteTimeout,
			ScheduledShutdown: serverScheduledShutdown,
			PeopleInfo:        peopleInfo,
		},
		DB: DBConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
		},
		Logger: LoggerConfig{
			Source: loggerSource,
		},
	}

	return config, nil
}
