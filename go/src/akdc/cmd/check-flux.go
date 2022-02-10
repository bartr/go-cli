/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fluxCmd represents the flux command
var checkFluxCmd = &cobra.Command{
	Use:   "flux",
	Short: "check flux status on each cluster",
	Long: `check flux status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check-flux called")
	},
}

func init() {
	checkCmd.AddCommand(checkFluxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fluxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fluxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
