/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"mycli/cmd"
	"os"
	"strconv"
)

type task struct {
	name string
	description string
	completed bool
}

func main() {
	cmd.Execute()
	tasks := make(map[int]task)
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
			if len(tasks) == 0 {
				fmt.Println("No hay tareas disponibles")
			}
			newTask := bufio.NewScanner(os.Stdin)
			fmt.Print("The new task is: ")
			fmt.Println(newTask.Scan())
		case 2:
		case 3:
		case 4:
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)
		}

	}
	
}
