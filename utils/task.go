package utils

import (
	"time"
)

type Task struct {
	TaskName string
	Completed bool
	Date string
}


type TaskList struct {
	list []Task
	Type string
}

func CreateTaskList() TaskList {
	return TaskList{}
}

func CreateTask(TaskName string) Task {
	task := Task {
		TaskName: TaskName,
		Completed: false,
		Date: time.Now().Format("2006-01-02"),
	}
	return task
}


func (t *TaskList) ReadTask(i int) Task {
	return t.list[i]
}

func (t *TaskList) GetTaskList() []Task {
	return t.list
}

func (t *TaskList) AddTask(task Task) *TaskList {
	t.list = append(t.list, task)
	return t
}

func (t *TaskList) DeleteTask(i int) *TaskList {
	t.list = append(t.list[:i],t.list[i+1:]... )
	return t
}

func (t *TaskList) EditTask(i int, text string) *TaskList {
	t.list[i].TaskName = text
	return t
}

func (t *TaskList) ToggleCompleteTask(i int) *TaskList{
	t.list[i].Completed = !t.list[i].Completed
	return t
}

