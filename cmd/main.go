package main

import (
	"cmd/main.go/configs"
	"cmd/main.go/logger"
	"cmd/main.go/repository"
	"cmd/main.go/server"
	"cmd/main.go/service"
	"cmd/main.go/transport"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger.Init(config.Logger.Source+"debug.log", config.Logger.Source+"info.log")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Настройка драйвера для postgres
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Настройка источника миграций
	source, err := (&file.File{}).Open("db/migrations")
	if err != nil {
		log.Fatal(err)
	}

	// Создание мигратора
	m, err := migrate.NewWithInstance("file", source, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	// Применение миграций
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	er := db.Ping()
	if er != nil {
		log.Fatal("db.Ping() возвращает ошибку: ", err)
	}
	logger.Info.Println("Миграции успешно применены!")

	myrepository := repository.NewRepository(db)
	myservice := service.NewService(myrepository, config)
	myhandler := transport.NewMyHandler(myservice)
	server := new(server.Server)

	if err := server.RunServer(config.Server.Port, myhandler.InitRoutes()); err != nil {
		log.Fatal("Server start error: ", err)
	}
	fmt.Scanln()
}
