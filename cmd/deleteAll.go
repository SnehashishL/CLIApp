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

// deleteAllCmd represents the deleteAll command
var deleteAllCmd = &cobra.Command{
	Use:   "deleteAll",
	Short: "Delete all tasks from DB",
	Long: `Delete all tasks from the todos DB. This command calls the REST API deleteTasks from the TodoApp and
	returns the status of the operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteAll called")
		fmt.Printf("\n\n    **** Deleting All tasks ****\n\n")
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/deleteAll/", nil)
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
	rootCmd.AddCommand(deleteAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
