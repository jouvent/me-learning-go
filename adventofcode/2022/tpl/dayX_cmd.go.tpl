/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package day0

import (
	"github.com/spf13/cobra"

	"aoc2022/day3"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Run the solution for the day one puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		input, part := flags(cmd, 3)
		day3.Process(input, part)

	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
