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

func (t *Task) Get() struct {
	ID           int64
	Name         string
	Description  string
	CreatedTime  time.Time
	IsDone       bool
	FinishedTime *time.Time
} {
	return struct {
		ID           int64
		Name         string
		Description  string
		CreatedTime  time.Time
		IsDone       bool
		FinishedTime *time.Time
	}{
		ID:           t.id,
		Name:         t.name,
		Description:  t.description,
		CreatedTime:  t.createdTime,
		IsDone:       t.isDone,
		FinishedTime: t.finishedTime,
	}
}
