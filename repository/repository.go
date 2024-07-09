package repository

import (
	"cmd/main.go/configs"
	"database/sql"
)

type repository struct {
	db     *sql.DB
	config *configs.Config
}

type Repository interface {
	WorkLogRepository
	UserRepository
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
