package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/4227078/ImageToTextTask
 */
type ImageToTextTask struct {
	Task
	Body      string                   `json:"body"`
	Phrase    bool                     `json:"phrase"`
	Case      bool                     `json:"case"`
	Numeric   int                      `json:"numeric"`
	Math      bool                     `json:"math"`
	MinLength int                      `json:"minLength"`
	MaxLength int                      `json:"maxLength"`
	Comment   string                   `json:"comment"`
	Timeout   time.Duration            `json:"-"`
	Response  *ImageToTextTaskResponse `json:"-"`
}

type ImageToTextTaskResponse struct {
	TaskResponse
	Solution *ImageToTextTaskSolution `json:"solution"`
}

type ImageToTextTaskSolution struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

func (t *ImageToTextTask) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &ImageToTextTaskResponse{}
}

func (t *ImageToTextTask) GetType() string {
	t.Type = "ImageToTextTask"
	return t.Type
}

func (t *ImageToTextTask) GetTimeout() time.Duration {
	var min = 5 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
