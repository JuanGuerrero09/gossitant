package utils

import (
	"testing"
)

var newTaskList = taskList{
	list: []task{
		{TaskName: "Tarea 1", Completed: false, Date: "2023-10-03"},
		{TaskName: "Tarea 2", Completed: true, Date: "2023-10-04"},
	},
}

func TestAddTask(t *testing.T) {
	newTask := task{
		TaskName:  "Tarea 3",
		Completed: false,
		Date:      "01-03-2022",
	}
	newTaskList.AddTask(newTask)
	got := newTaskList.ReadTask(2)
	want := taskList{
		list: []task{
			{TaskName: "Tarea 1", Completed: false, Date: "2023-10-03"},
			{TaskName: "Tarea 2", Completed: true, Date: "2023-10-04"},
			{TaskName: "Tarea 3", Completed: false, Date: "01-03-2022"},
		},
	}.list[2]
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDeleteTask(t *testing.T) {
	newTaskList.DeleteTask(0)
	got := newTaskList.GetTaskList()[0]
	want := taskList{
		list: []task{
			{TaskName: "Tarea 2", Completed: true, Date: "2023-10-04"},
		},
	}.list[0]
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}


func TestCompleteTask(t *testing.T) {
	taskList.ToggleCompleteTask(0)
	got := taskList.ReadTask(0)
	want := TaskList{
		list: []Task{
			{TaskName: "Tarea 2", Completed: false, Date: "2023-10-04"},
		},
	}.list[0]
	if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	