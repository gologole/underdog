package service

import (
	"cmd/main.go/logger"
	"cmd/main.go/models"
	"sort"
	"time"
)

type WorkLogService interface {
	GetWorkLogs(userID int, start, end time.Time) ([]*models.WorkLog, error)
	StartWork(userID, taskID int) error
	StopWork(userID, taskID int) error
}

func (s *Service) GetWorkLogs(userID int, start, end time.Time) ([]*models.WorkLog, error) {
	var wls []*models.WorkLog
	wls, err := s.r.GetWorkLogs(userID, start, end)
	if err != nil {
		logger.Debug.Printf("Failed to get work logs: %v", err)
		return nil, err
	}

	wl, er := sortByDuration(wls)
	if er != nil {
		logger.Debug.Printf("Будет смешно если здесь выпадет ошибка когда-нибудь: %v", err)
		return nil, err
	}
	return wl, nil
}

func parseDuration(d string) (time.Duration, error) {
	return time.ParseDuration(d)
}

func sortByDuration(logs []*models.WorkLog) ([]*models.WorkLog, error) {
	// Копируем оригинальный слайс, чтобы не изменять его напрямую
	sortedLogs := make([]*models.WorkLog, len(logs))
	copy(sortedLogs, logs)

	// Сортируем слайс
	sort.Slice(sortedLogs, func(i, j int) bool {
		durationI, err := parseDuration(sortedLogs[i].Duration)
		if err != nil {
			logger.Debug.Printf("Error parsing duration for ID %d: %v\n", sortedLogs[i].Id, err)
			return false
		}
		durationJ, err := parseDuration(sortedLogs[j].Duration)
		if err != nil {
			logger.Debug.Printf("Error parsing duration for ID %d: %v\n", sortedLogs[j].Id, err)
			return true
		}
		return durationI > durationJ
	})

	return sortedLogs, nil
}

func (s *Service) StartWork(userID, taskID int) error {
	return s.r.StartWork(userID, taskID)
}

func (s *Service) StopWork(userID, taskID int) error {
	return s.r.StopWork(userID, taskID)
}
