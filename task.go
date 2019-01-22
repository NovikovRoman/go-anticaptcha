package anticaptcha

import (
	"time"
)

type Tasker interface {
	Reset()
	SetTaskId(*CreateTaskResponse) int
	GetTaskId() int
	GetType() string
	GetTimeout() time.Duration
}

type Task struct {
	Type   string `json:"type"`
	TaskId int    `json:"taskId"`
}

func (t *Task) SetTaskId(resp *CreateTaskResponse) int {
	t.TaskId = resp.TaskId
	return t.TaskId
}

func (t *Task) GetTaskId() int {
	return t.TaskId
}

type CreateTaskResponse struct {
	ErrorResponse
	TaskId int `json:"taskId"`
}

type TaskResponse struct {
	ErrorResponse
	Status     string `json:"status"`
	Cost       string `json:"cost"`
	Ip         string `json:"ip"`
	CreateTime int    `json:"createTime"`
	EndTime    int    `json:"endTime"`
	SolveCount int    `json:"solveCount"`
}
