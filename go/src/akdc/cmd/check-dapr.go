/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// daprCmd represents the dapr command
var daprCmd = &cobra.Command{
	Use:   "dapr",
	Short: "check dapr status on each cluster",
	Long: `check dapr status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := getAllIps()

		for _, line := range lines {
			fmt.Println(line)
		}
	},
}

func init() {
	checkCmd.AddCommand(daprCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daprCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daprCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
