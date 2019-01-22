package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/417005605/GeeTestTaskProxyless+-+geetest.com
 */
type GeeTestTaskProxyless struct {
	Task
	WebsiteURL                string                        `json:"websiteURL"`
	Gt                        string                        `json:"gt"`
	Challenge                 string                        `json:"challenge"`
	GeetestApiServerSubdomain string                        `json:"geetestApiServerSubdomain"`
	Timeout                   time.Duration                 `json:"-"`
	Response                  *GeeTestTaskProxylessResponse `json:"-"`
}

type GeeTestTaskProxylessResponse struct {
	TaskResponse
	Solution *GeeTestTaskProxylessSolution `json:"solution"`
}

type GeeTestTaskProxylessSolution struct {
	Challenge string `json:"challenge"`
	Validate  string `json:"validate"`
	Seccode   string `json:"seccode"`
}

func (t *GeeTestTaskProxyless) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &GeeTestTaskProxylessResponse{}
}

func (t *GeeTestTaskProxyless) GetType() string {
	t.Type = "GeeTestTaskProxyless"
	return t.Type
}

func (t *GeeTestTaskProxyless) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
