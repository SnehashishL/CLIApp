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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to DB",
	Long: `Add a task to the todos DB using this command. This command calls the REST API addTasks from the TodoApp and
	returns the status of execution`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		taskid, _ := cmd.Flags().GetString("id")
		taskname, _ := cmd.Flags().GetString("name")

		fmt.Printf("\n\n    **** Adding task : %s ****\n\n", taskname)
		response, err := http.Post("http://localhost:8000/addTask/?taskid="+taskid+"&taskname="+taskname, "application/json; charset=utf-8", nil)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.PersistentFlags().String("id", "", "add task id and name")
	addCmd.PersistentFlags().String("name", "", "add task id and name")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
