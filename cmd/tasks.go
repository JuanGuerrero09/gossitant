/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"mycli/utils/task"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// newTaskCmd represents the newTask command
var newTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.CreateTaskList()
		for {
			fmt.Println("Select option")
			fmt.Println("1. Ver tareas")
			fmt.Println("2. Agregar tareas")
			fmt.Println("3. Editar tarea")
			fmt.Println("4. Marcar como completada")
			fmt.Println("5. Eliminar tarea")
			fmt.Println("6. Salir")
			fmt.Print("Opción: ")

			existentFile, err := os.Create("taskList.csv")
			if err != nil {
				panic(err)
			}

			

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			choice, _ := strconv.Atoi(scanner.Text())

			switch choice {
			case 1:
				if len(tasks.GetTaskList()) == 0 {
					fmt.Println("No hay tareas disponibles")
				}
				fmt.Println("The existent task are: ")
				for i, task := range tasks.GetTaskList() {
					fmt.Println(i+1, task.TaskName)
				}
				fmt.Println(tasks)
			case 2:
				fmt.Print("Agregue nueva tarea:...")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				newTask := task.CreateTask(scanner.Text())
				fmt.Println(newTask)
				tasks.AddTask(newTask)
			case 3:
				fmt.Println("Seleccione el índice de la tarea que desea editar")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				index, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("There was a problem with the index")
				}
				fmt.Println("Inserte el cambio de la tarea")
				scannerText := bufio.NewScanner(os.Stdin)
				scannerText.Scan()
				tasks.EditTask(index, scannerText.Text())
				fmt.Printf("Task number %v, changed to: %s\n", index, scannerText.Text())
			case 4:
				fmt.Println("Seleccione el índice de la tarea que desea marcar como completar")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				index, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("There was a problem with the index")
				}
				tasks.ToggleCompleteTask(index)
			case 5:
				fmt.Println("Guardando en txt....")
				newFile, err := os.Create("taskList.csv")
				if err != nil {
					panic(err)
				}

				for _, task := range tasks.GetTaskList() {
					strArr := make([]string, 0)
					strArr = append(strArr, task.TaskName, strconv.FormatBool(task.Completed), task.Date)
					str := strings.Join(strArr, ",")
					newFile.WriteString(str + "\n")
				}

				defer newFile.Close()
			case 6:
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
