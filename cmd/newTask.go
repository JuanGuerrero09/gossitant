/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"mycli/utils"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type task struct {
	taskDescription string
	completed bool
	date string
}



// newTaskCmd represents the newTask command
var newTaskCmd = &cobra.Command{
	Use:   "nt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("newTask called")
	tasks := utils.TaskList {}
	// lastId := 0
	fmt.Println("im here")
	
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
	},
}

func init() {
	rootCmd.AddCommand(newTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
