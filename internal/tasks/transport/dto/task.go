package dto

import "time"

type Task struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	CreatedTime  time.Time  `json:"created_time"`
	IsDone       bool       `json:"is_done"`
	FinishedTime *time.Time `json:"finished_time"`
}
