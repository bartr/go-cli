/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var checkSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "check setup status on each cluster",
	Long: `check setup status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check-setup called")
	},
}

func init() {
	checkCmd.AddCommand(checkSetupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
