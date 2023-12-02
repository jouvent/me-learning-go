/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"aoc2023/day2"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Run the solution for the day one puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		input, part := flags(cmd, 2)
		day2.Process(input, part)

	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
