package io

import (
	// "fmt"
	"os"
)



func SaveFile(){
	newFile, err := os.Create("taskList.csv")
	if err != nil {
		panic(err)
		}
	
	defer newFile.Close()


}

func LoadTasks() {
	
}

