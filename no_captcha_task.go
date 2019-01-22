package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/4227081/NoCaptchaTask+Google
 */
type NoCaptchaTask struct {
	Task
	WebsiteURL    string                 `json:"websiteURL"`
	WebsiteKey    string                 `json:"websiteKey"`
	WebsiteSToken string                 `json:"websiteSToken"`
	ProxyType     string                 `json:"proxyType"`
	ProxyAddress  string                 `json:"proxyAddress"`
	ProxyPort     string                 `json:"proxyPort"`
	ProxyLogin    string                 `json:"proxyLogin"`
	ProxyPassword string                 `json:"proxyPassword"`
	UserAgent     string                 `json:"userAgent"`
	Cookies       string                 `json:"cookies"`
	IsInvisible   bool                   `json:"isInvisible"`
	Timeout       time.Duration          `json:"-"`
	Response      *NoCaptchaTaskResponse `json:"-"`
}

type NoCaptchaTaskResponse struct {
	TaskResponse
	Solution *NoCaptchaTaskSolution `json:"solution"`
}

type NoCaptchaTaskSolution struct {
	GRecaptchaResponse    string `json:"gRecaptchaResponse"`
	GRecaptchaResponseMD5 string `json:"gRecaptchaResponseMD5"`
}

func (t *NoCaptchaTask) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &NoCaptchaTaskResponse{}
}

func (t *NoCaptchaTask) GetType() string {
	t.Type = "NoCaptchaTask"
	return t.Type
}

func (t *NoCaptchaTask) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
