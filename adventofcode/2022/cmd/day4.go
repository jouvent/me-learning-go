/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"aoc2022/day4"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Run the solution for the day one puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		input, part := flags(cmd, 4)
		day4.Process(input, part)

	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
