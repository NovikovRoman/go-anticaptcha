package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/9666604/NoCaptchaTaskProxyless+Google
 */
type NoCaptchaTaskProxyless struct {
	Task
	WebsiteURL    string                          `json:"websiteURL"`
	WebsiteKey    string                          `json:"websiteKey"`
	WebsiteSToken string                          `json:"websiteSToken"`
	IsInvisible   bool                            `json:"isInvisible"`
	Timeout       time.Duration                   `json:"-"`
	Response      *NoCaptchaTaskProxylessResponse `json:"-"`
}

type NoCaptchaTaskProxylessResponse struct {
	TaskResponse
	Solution *NoCaptchaTaskProxylessSolution `json:"solution"`
}

type NoCaptchaTaskProxylessSolution struct {
	GRecaptchaResponse    string `json:"gRecaptchaResponse"`
	GRecaptchaResponseMD5 string `json:"gRecaptchaResponseMD5"`
}

func (t *NoCaptchaTaskProxyless) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &NoCaptchaTaskProxylessResponse{}
}

func (t *NoCaptchaTaskProxyless) GetType() string {
	t.Type = "NoCaptchaTaskProxyless"
	return t.Type
}

func (t *NoCaptchaTaskProxyless) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
