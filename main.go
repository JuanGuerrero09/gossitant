/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"mycli/cmd"
	"mycli/utils"
	"os"
	"strconv"
	"time"
)






func main() {
	cmd.Execute()
	tasks := utils.TaskList {}
	// lastId := 0

	for {
		fmt.Println("Select option")
		fmt.Println("1. Ver tareas")
		fmt.Println("2. Agregar tareas")
		fmt.Println("3. Marcar como completada")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Print("Opción: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ :=strconv.Atoi(scanner.Text())
		
		switch choice {
		case 1:
			if len(tasks.GetTaskList()) == 0 {
				fmt.Println("No hay tareas disponibles")
			}
			fmt.Print("The new task is: ")
		case 2:
			fmt.Print("Agregue nueva tarea:...")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newTask := utils.Task {
				TaskDescription: scanner.Text(),
				Completed: false,
				Date: time.Now().Format("2006-01-02"),
			}
			fmt.Println(newTask)
			tasks.AddTask(newTask)


		case 3:
		case 4:
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)
		}

	}
	
}
