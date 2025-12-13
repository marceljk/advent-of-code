package year2025

import (
	"github.com/marceljk/advent-of-code/cmd/2025/day1"
	"github.com/marceljk/advent-of-code/cmd/2025/day10"
	"github.com/marceljk/advent-of-code/cmd/2025/day2"
	"github.com/marceljk/advent-of-code/cmd/2025/day3"
	"github.com/marceljk/advent-of-code/cmd/2025/day4"
	"github.com/marceljk/advent-of-code/cmd/2025/day5"
	"github.com/marceljk/advent-of-code/cmd/2025/day6"
	"github.com/marceljk/advent-of-code/cmd/2025/day7"
	"github.com/marceljk/advent-of-code/cmd/2025/day8"
	"github.com/marceljk/advent-of-code/cmd/2025/day9"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2025",
		Short: "Solutions for 2025",
		Long:  "Solutions for 2025.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	addSubcommands(cmd)
	return cmd
}

func addSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(day1.NewCmd())
	cmd.AddCommand(day2.NewCmd())
	cmd.AddCommand(day3.NewCmd())
	cmd.AddCommand(day4.NewCmd())
	cmd.AddCommand(day5.NewCmd())
	cmd.AddCommand(day6.NewCmd())
	cmd.AddCommand(day7.NewCmd())
	cmd.AddCommand(day8.NewCmd())
	cmd.AddCommand(day9.NewCmd())
	cmd.AddCommand(day10.NewCmd())
}
