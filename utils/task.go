package utils

import (
	"fmt"
	// "time"
)

type Task struct {
	TaskDescription string
	Completed bool
	Date string
}

type TaskList struct {
	list []Task
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
	return t
}

func (t *TaskList) EditTask(i int, text string) *TaskList {
	return t
}

func (t *TaskList) ToggleCompleteTask(i int) *TaskList{
	t.list[i].Completed = !t.list[i].Completed
	return t
}

func tasks(){
	fmt.Println("Hola mundo")
}

