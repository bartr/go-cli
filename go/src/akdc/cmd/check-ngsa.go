/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ngsaCmd represents the ngsa command
var checkNgsaCmd = &cobra.Command{
	Use:   "ngsa",
	Short: "check ngsa status on each cluster",
	Long: `check ngsa status on each cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check-ngsa called")
	},
}

func init() {
	checkCmd.AddCommand(checkNgsaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ngsaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ngsaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
