package domain

import "time"

type Task struct {
	id           int64
	name         string
	description  string
	createdTime  time.Time
	isDone       bool
	finishedTime *time.Time
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
		id:           id,
		name:         name,
		description:  description,
		createdTime:  createdTime,
		isDone:       isDone,
		finishedTime: finishedTime,
	}
}
