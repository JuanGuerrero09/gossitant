/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "mycli/utils"
	"github.com/spf13/cobra"
	"github.com/cli/browser"
)

// pmbaCmd represents the pmba command
var pmbaCmd = &cobra.Command{
	Use:   "pmba",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		urls := []string{"https://pkg.go.dev/os/exec", "https://pkg.go.dev/os/"}
		// utils.Open(urls)
		browser.OpenURL(urls[0])
		fmt.Println("pmba called")
	},
}

func init() {
	rootCmd.AddCommand(pmbaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pmbaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pmbaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
