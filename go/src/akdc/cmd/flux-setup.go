/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reconcileCmd represents the reconcile command
var fluxSetupCmd = &cobra.Command{
	Use:   "reconcile",
	Short: "setup flux on each cluster",
	Long: `setup flux on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flux-setup called")
	},
}

func init() {
	fluxCmd.AddCommand(fluxSetupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reconcileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reconcileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
