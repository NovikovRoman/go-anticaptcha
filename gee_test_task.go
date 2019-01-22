package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/416907286/GeeTestTask+-+geetest.com
 */
type GeeTestTask struct {
	Task
	WebsiteURL                string               `json:"websiteURL"`
	Gt                        string               `json:"gt"`
	Challenge                 string               `json:"challenge"`
	GeetestApiServerSubdomain string               `json:"geetestApiServerSubdomain"`
	ProxyType                 string               `json:"proxyType"`
	ProxyAddress              string               `json:"proxyAddress"`
	ProxyPort                 string               `json:"proxyPort"`
	ProxyLogin                string               `json:"proxyLogin"`
	ProxyPassword             string               `json:"proxyPassword"`
	UserAgent                 string               `json:"userAgent"`
	Cookies                   string               `json:"cookies"`
	Timeout                   time.Duration        `json:"-"`
	Response                  *GeeTestTaskResponse `json:"-"`
}

type GeeTestTaskResponse struct {
	TaskResponse
	Solution *GeeTestTaskSolution `json:"solution"`
}

type GeeTestTaskSolution struct {
	Challenge string `json:"challenge"`
	Validate  string `json:"validate"`
	Seccode   string `json:"seccode"`
}

func (t *GeeTestTask) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &GeeTestTaskResponse{}
}

func (t *GeeTestTask) GetType() string {
	t.Type = "GeeTestTask"
	return t.Type
}

func (t *GeeTestTask) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
