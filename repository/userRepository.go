package repository

import (
	"cmd/main.go/logger"
	"cmd/main.go/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRepository interface {
	Create(user models.People) error
	GetByID(id int) (*models.People, error)
	FilterPeople(filters models.People, pagination models.Pagination) ([]models.People, error)
	Update(user *models.People) error
	Delete(user *models.People) error
}

func (r *repository) Create(user models.People) error {
	query := `INSERT INTO people (surname, name, patronymic, address, passport_number) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, user.Surname, user.Name, user.Patronymic, user.Address, user.PassportNumber).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByID(id int) (*models.People, error) {
	query := `SELECT id, surname, name, patronymic, address, passport_number FROM people WHERE id = $1`
	var user models.People
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.PassportNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(user *models.People) error {
	query := `UPDATE people SET surname = $1, name = $2, patronymic = $3, address = $4, passport_number = $5 WHERE id = $6`
	_, err := r.db.Exec(query, user.Surname, user.Name, user.Patronymic, user.Address, user.PassportNumber, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(user *models.People) error {
	query := `DELETE FROM people WHERE surname = $1 AND name = $2`
	_, err := r.db.Exec(query, user.Surname, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FilterPeople(filters models.People, pagination models.Pagination) ([]models.People, error) {
	var people []models.People

	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT * FROM people WHERE 1=1")

	if filters.Surname != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND surname = '%s'", filters.Surname))
	}
	if filters.Name != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND name = '%s'", filters.Name))
	}
	if filters.Patronymic != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND patronymic = '%s'", filters.Patronymic))
	}
	if filters.Address != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND address = '%s'", filters.Address))
	}
	if filters.PassportNumber != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND passport_number = '%s'", filters.PassportNumber))
	}

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder.WriteString(fmt.Sprintf(" OFFSET %d LIMIT %d", offset, pagination.Limit))

	query := queryBuilder.String()
	logger.Info.Printf("Executing query: %s", query)

	rows, err := r.db.Query(query)
	if err != nil {
		logger.Info.Printf("Failed to execute query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person models.People
		if err := rows.Scan(&person.ID, &person.Surname, &person.Name, &person.Patronymic, &person.Address, &person.PassportNumber); err != nil {
			logger.Info.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		people = append(people, person)
	}

	if err := rows.Err(); err != nil {
		logger.Info.Printf("Rows iteration error: %v", err)
		return nil, err
	}

	return people, nil
}
