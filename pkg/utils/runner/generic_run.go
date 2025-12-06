package runner

import (
	"fmt"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/spf13/cobra"
)

type TaskFunc func(parseInput string) error

func GenericRunE(task1, task2 TaskFunc) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		model, err := globalflags.ParseInput(cmd)
		if err != nil {
			return fmt.Errorf("failed parsing model: %w", err)
		}

		switch model.Task {
		case 2:
			if task2 == nil {
				return fmt.Errorf("task 2 is not implemented yet")
			}
			return task2(model.FileContent)
		default:
			if task1 == nil {
				return fmt.Errorf("task 1 is not implemented yet")
			}
			return task1(model.FileContent)
		}
	}
}
