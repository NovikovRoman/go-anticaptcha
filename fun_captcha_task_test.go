package anticaptcha

import (
	"testing"
	"time"
)

func TestFunCaptchaTask_GetType(t *testing.T) {
	task := &FunCaptchaTask{}
	task.Type = "test"
	if task.GetType() == "test" {
		t.Errorf("Type неверный %s", task.GetType())
	}
}

func TestFunCaptchaTask_Reset(t *testing.T) {
	task := &FunCaptchaTask{}
	resp := &CreateTaskResponse{TaskId: 111}
	task.SetTaskId(resp)
	task.Type = "test"
	task.Response = &FunCaptchaTaskResponse{}
	task.Response.Solution = &FunCaptchaTaskSolution{}
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

func TestFunCaptchaTask_GetTimeout(t *testing.T) {
	var min = time.Duration(time.Second * 10)
	task := &FunCaptchaTask{
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
