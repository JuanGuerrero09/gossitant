package utils

import (
	"testing"
)

var taskList = TaskList{
	list: []Task{
		{TaskDescription: "Tarea 1", Completed: false, Date: "2023-10-03"},
		{TaskDescription: "Tarea 2", Completed: true, Date: "2023-10-04"},
	},
}

func TestAddTask(t *testing.T) {
	task := Task{
		TaskDescription: "Tarea 3",
		Completed:       false,
		Date:            "01-03-2022",
	}
	taskList.AddTask(task)
	got := taskList.ReadTask(2)
	want := TaskList{
		list: []Task{
			{TaskDescription: "Tarea 1", Completed: false, Date: "2023-10-03"},
			{TaskDescription: "Tarea 2", Completed: true, Date: "2023-10-04"},
			{
				TaskDescription: "Tarea 3",
				Completed:       false,
				Date:            "01-03-2022",
			},
		},
	}.list[2]
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDeleteTask(t *testing.T) {
	taskList.DeleteTask(0)
	got := taskList.GetTaskList()[0]
	want := TaskList {
		list: []Task {
			{TaskDescription: "Tarea 2", Completed: true, Date: "2023-10-04"},
		},
	}.list[0]
	
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
