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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from DB",
	Long: `Delete a specific task from the todos DB. This command calls the REST API deleteTask from the TodoApp and
	returns the status of the operation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		id, _ := cmd.Flags().GetString("id")
		fmt.Printf("\n\n    **** Deleting task : %s ****\n\n", id)
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/deleteTask/"+id, nil)
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
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	deleteCmd.PersistentFlags().String("id", "", "delete task by id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
