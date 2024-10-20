/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/todo_app/extract"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "tasks delete <id>",
	Long: `delete a task from the task list`,
	Run: func(cmd *cobra.Command, args []string) {
    id := args[0]
    tasks := extract.ReadCsv("test.csv")

    f, err := os.Create("test.csv")
    defer f.Close()

    if err != nil {
      panic(err)
    }

    idExists := false

    for i, line := range tasks{
      if line[0] == id{
        tasks = append(tasks[:i], tasks[i+1:]...)
        idExists = true
      }
    }

    if !idExists{
      fmt.Printf("Attention, la tache d'id %s n'existe pas\n", id)
    }

    w := csv.NewWriter(f)
    err = w.WriteAll(tasks)

    if err != nil {
      panic(err)
    }

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
