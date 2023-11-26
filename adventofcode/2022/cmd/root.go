/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of code 2022",
	Long: `My completion of the Advent of Code 2022 puzzles
for more info, visit: https://adventofcode.com/2022/about
examples and usage:

    aoc day16 --part 2 --input input.txt
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int("part", 1, "part 1 or 2 of the challenge")
	rootCmd.PersistentFlags().String("input", "", "path to the input file (default is ./input_dayX.txt)")
}
