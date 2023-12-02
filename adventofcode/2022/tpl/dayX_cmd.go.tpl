/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"aoc2022/dayX"
)

// dayXCmd represents the dayX command
var dayXCmd = &cobra.Command{
	Use:   "dayX",
	Short: "Run the solution for the day one puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		input, part := flags(cmd, X)
		dayX.Process(input, part)

	},
}

func init() {
	rootCmd.AddCommand(dayXCmd)
}
