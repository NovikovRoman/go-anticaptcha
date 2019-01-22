package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/410714125/SquareNetTextTask
 */
type SquareNetTextTask struct {
	Task
	Body         string                     `json:"body"`
	ObjectName   string                     `json:"objectName"`
	RowsCount    int                        `json:"rowsCount"`
	ColumnsCount int                        `json:"columnsCount"`
	Timeout      time.Duration              `json:"-"`
	Response     *SquareNetTextTaskResponse `json:"-"`
}

type SquareNetTextTaskResponse struct {
	TaskResponse
	Solution *SquareNetTextTaskSolution `json:"solution"`
}

type SquareNetTextTaskSolution struct {
	CellNumbers []int `json:"cellNumbers"`
}

func (t *SquareNetTextTask) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &SquareNetTextTaskResponse{}
}

func (t *SquareNetTextTask) GetType() string {
	t.Type = "SquareNetTextTask"
	return t.Type
}

func (t *SquareNetTextTask) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
