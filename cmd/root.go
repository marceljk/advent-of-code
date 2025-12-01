package cmd

import (
	"os"

	year2025 "github.com/marceljk/advent-of-code/cmd/2025"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent-of-code",
	Short: "CLI to run advent of code solutions",
	Long:  `CLI to run advent of code solutions`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	addSubcommands(rootCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(year2025.NewCmd())
}
