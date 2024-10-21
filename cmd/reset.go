/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func resetFiles(){
  idFile, err := os.OpenFile("currentID.txt", os.O_WRONLY, 0644)
  defer idFile.Close()
  if err != nil {
    panic(err)
  }

  idFile.WriteString("0") 
  fmt.Println("tasks ID réinitialisé à 1 avec succès")

  tasksFile, err := os.Create("test.csv")
  defer idFile.Close()
  if err != nil {
    panic(err)
  }

  tasks := []string{"Id", "Description", "Created", "Completed"}
  w := csv.NewWriter(tasksFile)
  w.Write(tasks)
  w.Flush()
  fmt.Println("taches supprimées avec succès")
}

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    var confirmation string
    fmt.Print("Attention êtes vous sur de vouloir remettre a 0 toute votre liste de taches ? [y|n] ")
    fmt.Scanln(&confirmation)
    if confirmation == "y"{
      resetFiles()
    }
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
