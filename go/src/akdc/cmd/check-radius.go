/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// radiusCmd represents the radius command
var checkRadiusCmd = &cobra.Command{
	Use:   "radius",
	Short: "check radius status on each cluster",
	Long: `check radius status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check-radius called")
	},
}

func init() {
	checkCmd.AddCommand(checkRadiusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// radiusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// radiusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
