package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func flags(cmd *cobra.Command, day int) (string, int) {
	part, err := cmd.Flags().GetInt("part")
	if err != nil {
		panic(err)
	}
	input, err := cmd.Flags().GetString("input")
	if err != nil {
		panic(err)
	}
	if input == "" {
		input = fmt.Sprintf("./input_day%d.txt", day)
	}
	return input, part
}
