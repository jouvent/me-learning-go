/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"aoc2022/day1"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Run the solution for the day one puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		input, part := flags(cmd, 1)
		day1.Process(input, part)

	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
