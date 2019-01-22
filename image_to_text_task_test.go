package anticaptcha

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestImageToTextTask_GetType(t *testing.T) {
	task := &ImageToTextTask{}
	task.Type = "test"
	if task.GetType() == "test" {
		t.Errorf("Type неверный %s", task.GetType())
	}
}

func TestImageToTextTask_Reset(t *testing.T) {
	task := &ImageToTextTask{}
	resp := &CreateTaskResponse{TaskId: 111}
	task.SetTaskId(resp)
	task.Type = "test"
	task.Response = &ImageToTextTaskResponse{}
	task.Response.Solution = &ImageToTextTaskSolution{}
	task.Reset()
	if task.GetTaskId() != 0 {
		t.Errorf("TaskId не сброшен TaskId: %d", task.GetTaskId())
	}
	if task.Response.Solution != nil {
		t.Errorf("Solution не сброшен %p", task.Response.Solution)
	}
	if task.GetType() == "test" {
		t.Errorf("Type неверный %s", task.GetType())
	}
}

func TestImageToTextTask_GetTaskResult(t *testing.T) {
	reader, err := os.Open("./captcha.gif")
	if err != nil {
		t.Fatalf("Error open file %s", err)
	}
	defer func() {
		if rerr := reader.Close(); rerr != nil {
			err = rerr
		}
	}()
	img, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("Error read file %s", err)
	}
	imageTextTask := &ImageToTextTask{
		Body:    base64.StdEncoding.EncodeToString(img),
		Numeric: 1,
	}
	ac := Client{
		ClientKey:    os.Getenv("API_KEY"),
		LanguagePool: "en",
	}
	resTask, err := ac.SendImageToTextTask(imageTextTask)
	if err != nil {
		t.Fatalf("Error task %s", err)
	}
	if resTask.Solution.Text != "595040" {
		t.Fatalf("Captcha is unsolved %s", err)
	}
	resTaskJson, err := json.Marshal(&resTask)
	if err != nil {
		t.Fatalf("Error marshal %s", err)
	}
	t.Logf("%s", string(resTaskJson))
}

func TestImageToTextTask_GetTimeout(t *testing.T) {
	var min = time.Duration(time.Second * 5)
	task := &ImageToTextTask{
		Timeout: time.Second * 1,
	}
	if task.GetTimeout() < min {
		t.Errorf("Timeout не должен быть меньше %s", min)
	}
	t.Logf("Timeout: %s", task.GetTimeout())
	task.Timeout = min * 2
	if task.GetTimeout() < min {
		t.Errorf("Timeout не должен быть меньше %s", min)
	}
	t.Logf("Timeout: %s", task.GetTimeout())
}
