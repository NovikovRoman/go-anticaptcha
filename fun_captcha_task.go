package anticaptcha

import (
	"time"
)

/**
Документация https://anticaptcha.atlassian.net/wiki/spaces/API/pages/65634354/FunCaptchaTask+-+funcaptcha.com
 */
type FunCaptchaTask struct {
	Task
	FuncaptchaApiJSSubdomain string                  `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string                  `json:"websitePublicKey"`
	ProxyType                string                  `json:"proxyType"`
	ProxyAddress             string                  `json:"proxyAddress"`
	ProxyPort                string                  `json:"proxyPort"`
	ProxyLogin               string                  `json:"proxyLogin"`
	ProxyPassword            string                  `json:"proxyPassword"`
	UserAgent                string                  `json:"userAgent"`
	Cookies                  string                  `json:"cookies"`
	Timeout                  time.Duration           `json:"-"`
	Response                 *FunCaptchaTaskResponse `json:"-"`
}

type FunCaptchaTaskResponse struct {
	TaskResponse
	Solution *FunCaptchaTaskSolution `json:"solution"`
}

type FunCaptchaTaskSolution struct {
	Token string `json:"token"`
}

func (t *FunCaptchaTask) Reset() {
	t.GetType()
	t.TaskId = 0
	t.Response = &FunCaptchaTaskResponse{}
}

func (t *FunCaptchaTask) GetType() string {
	t.Type = "FunCaptchaTask"
	return t.Type
}

func (t *FunCaptchaTask) GetTimeout() time.Duration {
	var min = 10 * time.Second
	if t.Timeout < min {
		t.Timeout = min
	}
	return t.Timeout
}
