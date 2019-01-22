package anticaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const Host = "api.anti-captcha.com"
const Scheme = "https"

type Client struct {
	ClientKey    string `json:"clientKey"`
	LanguagePool string `json:"languagePool"`
	CallbackUrl  string `json:"callbackUrl"`
	SoftId       int    `json:"softId"`
}

type ErrorResponse struct {
	ErrorId          int    `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

type BalanceResponse struct {
	ErrorResponse
	Balance float32 `json:"balance"`
}

type QueueStatsResponse struct {
	Waiting int     `json:"waiting"`
	Load    float32 `json:"load"`
	Bid     float32 `json:"bid"`
	Speed   float32 `json:"speed"`
	Total   int     `json:"total"`
}

type ReportResponse struct {
	ErrorId int     `json:"errorId"`
	Status  float32 `json:"status"`
}

func (c *Client) GetBalance() (*BalanceResponse, error) {
	reqBody := map[string]interface{}{
		"clientKey": c.ClientKey,
	}
	responseData := &BalanceResponse{}
	if err := c.request(&reqBody, "/getBalance", responseData); err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *Client) ReportIncorrectImageCaptcha(taskId int) (*ReportResponse, error) {
	reqBody := map[string]interface{}{
		"clientKey": c.ClientKey,
		"taskId":    taskId,
	}
	responseData := &ReportResponse{}
	if err := c.request(&reqBody, "/reportIncorrectImageCaptcha", responseData); err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *Client) GetQueueStats(queueId int) (*QueueStatsResponse, error) {
	reqBody := map[string]interface{}{
		"queueId": queueId,
	}
	responseData := &QueueStatsResponse{}
	if err := c.request(&reqBody, "/getQueueStats ", responseData); err != nil {
		return nil, err
	}
	return responseData, nil
}

// ImageToTextTask
func (c *Client) SendImageToTextTask(task *ImageToTextTask) (*ImageToTextTaskResponse, error) {
	task.Reset()
	reqBody := struct {
		Client
		Task *ImageToTextTask `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// FunCaptchaTask
func (c *Client) SendFunCaptchaTask(task *FunCaptchaTask) (*FunCaptchaTaskResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &FunCaptchaTaskResponse{}
	reqBody := struct {
		Client
		Task *FunCaptchaTask `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// FunCaptchaTaskProxyless
func (c *Client) SendFunCaptchaTaskProxyless(task *FunCaptchaTaskProxyless) (*FunCaptchaTaskProxylessResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &FunCaptchaTaskProxylessResponse{}
	reqBody := struct {
		Client
		Task *FunCaptchaTaskProxyless `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// GeeTestTask
func (c *Client) SendGeeTestTask(task *GeeTestTask) (*GeeTestTaskResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &GeeTestTaskResponse{}
	reqBody := struct {
		Client
		Task *GeeTestTask `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// GeeTestTaskProxyless
func (c *Client) SendGeeTestTaskProxyless(task *GeeTestTaskProxyless) (*GeeTestTaskProxylessResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &GeeTestTaskProxylessResponse{}
	reqBody := struct {
		Client
		Task *GeeTestTaskProxyless `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// NoCaptchaTask
func (c *Client) SendNoCaptchaTask(task *NoCaptchaTask) (*NoCaptchaTaskResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &NoCaptchaTaskResponse{}
	reqBody := struct {
		Client
		Task *NoCaptchaTask `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// NoCaptchaTaskProxyless
func (c *Client) SendNoCaptchaTaskProxyless(task *NoCaptchaTaskProxyless) (*NoCaptchaTaskProxylessResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &NoCaptchaTaskProxylessResponse{}
	reqBody := struct {
		Client
		Task *NoCaptchaTaskProxyless `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

// SquareNetTextTask
func (c *Client) SendSquareNetTextTask(task *SquareNetTextTask) (*SquareNetTextTaskResponse, error) {
	task.GetType()
	task.TaskId = 0
	task.Response = &SquareNetTextTaskResponse{}
	reqBody := struct {
		Client
		Task *SquareNetTextTask `json:"task"`
	}{
		Client: *c,
		Task:   task,
	}
	if err := c.send(task, reqBody); err != nil {
		return nil, err
	}
	reqTaskBody := c.createTaskBody(task)
	for task.Response.Status == "processing" || task.Response.Status == "" {
		time.Sleep(task.GetTimeout())
		if err := c.request(&reqTaskBody, "/getTaskResult", &task.Response); err != nil {
			return nil, err
		}
	}
	return task.Response, nil
}

func (c *Client) send(tasker Tasker, reqBody interface{}) error {
	responseData := &CreateTaskResponse{}
	if err := c.request(&reqBody, "/createTask", responseData); err != nil {
		return err
	}
	tasker.SetTaskId(responseData)
	if responseData.ErrorId == 0 {
		return nil
	}
	return fmt.Errorf("id: %d code: %s error: %s",
		responseData.ErrorId, responseData.ErrorCode, responseData.ErrorDescription,
	)
}

func (c *Client) createTaskBody(tasker Tasker) interface{} {
	return map[string]interface{}{
		"clientKey": c.ClientKey,
		"taskId":    tasker.GetTaskId(),
	}
}

func (c *Client) request(req interface{}, path string, responseData interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	requestUrl := &url.URL{Host: Host, Scheme: Scheme, Path: path}
	resp, err := http.Post(requestUrl.String(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err := json.NewDecoder(resp.Body).Decode(responseData); err != nil {
		return err
	}
	return err
}
