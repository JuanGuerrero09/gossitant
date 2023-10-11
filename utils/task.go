package utils

import (
	"time"
)

type task struct {
	TaskName  string
	Completed bool
	Date      string
}

type taskList struct {
	list []task
	Type string
}

func CreateTaskList() taskList {
	return taskList{}
}

func CreateTask(TaskName string) task {
	task := task{
		TaskName:  TaskName,
		Completed: false,
		Date:      time.Now().Format("2006-01-02"),
	}
	return task
}

func (t *taskList) ReadTask(i int) task {
	return t.list[i]
}

func (t *taskList) GetTaskList() []task {
	return t.list
}

func (t *taskList) AddTask(tk task) *taskList {
	t.list = append(t.list, tk)
	return t
}

func (t *taskList) DeleteTask(i int) *taskList {
	t.list = append(t.list[:i], t.list[i+1:]...)
	return t
}

func (t *taskList) EditTask(i int, text string) *taskList {
	t.list[i].TaskName = text
	return t
}

func (t *taskList) ToggleCompleteTask(i int) *taskList {
	t.list[i].Completed = !t.list[i].Completed
	return t
}
