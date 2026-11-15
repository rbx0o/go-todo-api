package tasks

import "time"

type Task struct {
	id           int64
	name         string
	description  string
	createdTime  time.Time
	isDone       bool
	finishedTime *time.Time
}
