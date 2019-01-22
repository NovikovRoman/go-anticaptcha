package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/327024659/FunCaptchaTaskProxyless+-+funcaptcha
 */
type FunCaptchaTaskProxyless struct {
	Task
	FuncaptchaApiJSSubdomain string                           `json:"funcaptchaApiJSSubdomain"`
	WebsiteURL               string                           `json:"websiteURL"`
	WebsitePublicKey         string                           `json:"websitePublicKey"`
	Timeout                  time.Duration                    `json:"-"`
	Response                 *FunCaptchaTaskProxylessResponse `json:"-"`
}

type FunCaptchaTaskProxylessResponse struct {
	TaskResponse
	Solution *FunCaptchaTaskProxylessSolution `json:"solution"`
}

type FunCaptchaTaskProxylessSolution struct {
	Token string `json:"token"`
}

func (t *FunCaptchaTaskProxyless) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &FunCaptchaTaskProxylessResponse{}
}

func (t *FunCaptchaTaskProxyless) GetType() string {
	t.Type = "FunCaptchaTaskProxyless"
	return t.Type
}

func (t *FunCaptchaTaskProxyless) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
