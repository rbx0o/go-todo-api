package domain

import "time"

type Task struct {
	ID           int64
	Name         string
	Description  string
	CreatedTime  time.Time
	IsDone       bool
	FinishedTime *time.Time
}

func NewTask(
	id int64,
	name string,
	description string,
	createdTime time.Time,
	isDone bool,
	finishedTime *time.Time,
) *Task {
	return &Task{
		ID:           id,
		Name:         name,
		Description:  description,
		CreatedTime:  createdTime,
		IsDone:       isDone,
		FinishedTime: finishedTime,
	}
}
