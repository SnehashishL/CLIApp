/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task in DB",
	Long: `Updates a task in the todos DB. This command calls the REST API updateTask from the TodoApp and
	returns the status of the operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		id, _ := cmd.Flags().GetString("id")
		newTask, _ := cmd.Flags().GetString("name")
		fmt.Printf("\n\n    **** Updating task %s with new name : %s****\n\n", id, newTask)
		req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/updateTask/"+id+"?taskname="+newTask, nil)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	updateCmd.PersistentFlags().String("id", "", "update task by id")
	updateCmd.PersistentFlags().String("name", "", "update task by id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
