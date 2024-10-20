/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"

	"github.com/todo_app/extract"
)

func listTasks(filePath string, showAll bool){
  tasks := extract.ReadCsv(filePath)
  w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
  for i, line := range tasks{
    if !(line[3] == "true" && !showAll){
      for j, val := range line{
        if i > 0 && j==2{
          fmt.Fprintf(w, "%s\t", timediff.TimeDiff(extract.StrToTime(val)))
        } else if !(j == 3 && !showAll){
          fmt.Fprintf(w, "%s\t", val)
        }
      }
      fmt.Fprint(w,"\n")
    }
  }
  w.Flush()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    showAll, _ := cmd.Flags().GetBool("all")
    listTasks("test.csv", showAll)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  listCmd.PersistentFlags().BoolP("all", "a", false, "show completed tasks as well")
}
