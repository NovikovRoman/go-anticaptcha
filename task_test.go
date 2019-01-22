package anticaptcha

import "testing"

func TestTask_TaskId(t *testing.T) {
	taskId := 111
	task := Task{
		TaskId: 222,
	}
	taskResponse := &CreateTaskResponse{TaskId: taskId}
	task.SetTaskId(taskResponse)
	if task.GetTaskId() != taskId {
		t.Errorf("TaskId не совпадает с %d TaskId: %d", taskId, task.GetTaskId())
	}
}
