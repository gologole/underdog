package service

import (
	"cmd/main.go/logger"
	"cmd/main.go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserService interface {
	CreateUser(user *models.People) error
	GetUsersListByParams(user *models.People, pagination models.Pagination) []models.People
	UpdateUser(user *models.People) error
	DeleteUser(id int) error
}

func (s *Service) CreateUser(user *models.People) error {
	personInfo, err := s.getPersonInfo(user.PassportNumber)
	if err != nil {
		return fmt.Errorf("failed to get person info: %v", err)
	}

	user.Surname = personInfo.Surname
	user.Name = personInfo.Name
	user.Patronymic = personInfo.Patronymic
	user.Address = personInfo.Address

	err = s.r.Create(*user)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	return nil
}

func (s *Service) GetUsersListByParams(user *models.People, pagination models.Pagination) ([]models.People, error) {
	users, err := s.r.FilterPeople(*user, pagination)
	if err != nil {
		logger.Info.Printf("Failed to filter people: %v", err)
		return nil, err
	}
	return users, nil
}

func (s *Service) getPersonInfo(passport string) (*models.People, error) {
	parts := strings.Split(passport, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid passport format")
	}

	passportSerie := parts[0]
	passportNumber := parts[1]

	url := fmt.Sprintf("%s/info?passportSerie=%s&passportNumber=%s", s.config.Server.PeopleInfo, passportSerie, passportNumber)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var personInfo models.People
	if err := json.NewDecoder(resp.Body).Decode(&personInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return &personInfo, nil
}

func (s *Service) UpdateUser(user *models.People) error {
	return s.r.Update(user)
}

func (s *Service) DeleteUser(user *models.People) error {
	return s.r.Delete(user)
}
