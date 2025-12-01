package globalflags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	inputFileFlag      = "file"
	inputFileFlagShort = "f"
	taskFlag           = "task"
)

type Model struct {
	FileContent string
	Task        uint
}

func ConfigureFlags(cmd *cobra.Command) {
	cmd.Flags().Uint(taskFlag, 1, "Task (Valid input: 1, 2)")
	cmd.Flags().StringP(inputFileFlag, inputFileFlagShort, "", "Path to input file.")
	cmd.MarkFlagRequired(inputFileFlag)
}

func ParseInput(cmd *cobra.Command) (*Model, error) {
	// Read file path from flag
	filePath, err := cmd.Flags().GetString(inputFileFlag)
	if err != nil {
		return nil, fmt.Errorf("can not read flag %q: %w", inputFileFlag, err)
	}

	// Read file from file path
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("can not find or read file from path %q: %w", filePath, err)
	}

	// Read task ID
	taskId, err := cmd.Flags().GetUint(taskFlag)
	if err != nil {
		return nil, fmt.Errorf("can not read task from flag %q: %w", taskId, err)
	}

	input := &Model{
		FileContent: string(content),
		Task:        taskId,
	}
	return input, nil
}
