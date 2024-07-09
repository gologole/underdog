package repository

import (
	"cmd/main.go/models"
	"time"
)

type WorkLogRepository interface {
	LogWork(userID, taskID int, duration time.Duration) error
	GetWorkLogs(userID int, start, end time.Time) ([]*models.WorkLog, error)
	StartWork(userID, taskID int) error
	StopWork(userID, taskID int) error
}

func (r *repository) LogWork(userID, taskID int, duration time.Duration) error {
	query := `INSERT INTO work_logs (user_id, task_id, start_time, end_time, duration) 
	          VALUES ($1, $2, $3, $4, $5)`
	startTime := time.Now()
	endTime := startTime.Add(duration)
	_, err := r.db.Exec(query, userID, taskID, startTime, endTime, duration.String())
	return err
}

// сортировка в сервис-слое
func (r *repository) GetWorkLogs(userID int, start, end time.Time) ([]*models.WorkLog, error) {
	query := `SELECT id, user_id, task_id, start_time, end_time, duration 
	          FROM work_logs 
	          WHERE user_id = $1 AND start_time >= $2 AND end_time <= $3`
	rows, err := r.db.Query(query, userID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workLogs []*models.WorkLog
	for rows.Next() {
		var workLog models.WorkLog
		err := rows.Scan(&workLog.Id, &workLog.UserID, &workLog.TaskID, &workLog.StartTime, &workLog.EndTime, &workLog.Duration)
		if err != nil {
			return nil, err
		}
		workLogs = append(workLogs, &workLog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workLogs, nil
}

func (r *repository) StartWork(userID, taskID int) error {
	query := `INSERT INTO work_logs (user_id, task_id, start_time) 
	          VALUES ($1, $2, $3)`
	startTime := time.Now()
	_, err := r.db.Exec(query, userID, taskID, startTime)
	return err
}

func (r *repository) StopWork(userID, taskID int) error {
	query := `UPDATE work_logs 
	          SET end_time = $1, duration = $2 
	          WHERE user_id = $3 AND task_id = $4 AND end_time IS NULL 
	          ORDER BY start_time DESC 
	          LIMIT 1`
	endTime := time.Now()
	var startTime time.Time

	err := r.db.QueryRow(`SELECT start_time FROM work_logs 
	                      WHERE user_id = $1 AND task_id = $2 AND end_time IS NULL 
	                      ORDER BY start_time DESC 
	                      LIMIT 1`, userID, taskID).Scan(&startTime)
	if err != nil {
		return err
	}

	duration := endTime.Sub(startTime)
	_, err = r.db.Exec(query, endTime, duration.String(), userID, taskID)
	return err
}
