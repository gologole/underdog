package models

import "time"

type People struct {
	ID             int    `json:"id"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address"`
	PassportNumber string `json:"passportNumber"`
}

type WorkLog struct {
	Id        int       `json:"id"`
	UserID    int       `json:"user_id" example:"1"`
	TaskID    int       `json:"task_id" example:"100"`
	StartTime time.Time `json:"start_time" example:"2023-01-01T15:04:05Z"`
	EndTime   time.Time `json:"end_time" example:"2023-01-01T17:04:05Z"`
	Duration  string    `json:"duration"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsActive    bool   `json:"isactive"`
	WorkLog     string `json:"workLog"` //в стрингу(json.Marshal) структуру когда в бд и демаршалить при чтении
}

type CreateUserRequest struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
