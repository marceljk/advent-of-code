package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils"
	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/spf13/cobra"
)

type Battery = uint
type Bank []Battery

func (b Bank) GetLargestJoltage(amountBatteries uint) (uint64, error) {
	if len(b) == 0 {
		return 0, fmt.Errorf("bank has no values")
	}
	activeBatteries := []uint{}
	bankSize := len(b)
	startIdx := 0
	endIdx := bankSize + 1 - int(amountBatteries) // end index is size with space of the amount of batteries
	firstBattery, startIdx := utils.FindLargestValue(b[startIdx:endIdx])
	endIdx++
	activeBatteries = append(activeBatteries, firstBattery)
	for range amountBatteries - 1 {
		nextBattery, nextIdx := utils.FindLargestValue(b[startIdx+1 : endIdx])
		startIdx = startIdx + nextIdx + 1
		activeBatteries = append(activeBatteries, nextBattery)
		endIdx++
	}
	sum := utils.ConcateUint(activeBatteries...)
	return sum, nil
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "3",
		Short: "Solution for 2025 day 3",
		Long:  "Solution for 2025 day 3.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			model, err := globalflags.ParseInput(cmd)
			if err != nil {
				return fmt.Errorf("failed parsing model: %w", err)
			}

			switch model.Task {
			case 2:
				return task2(*model)
			default:
				return task1(*model)
			}
		},
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(model globalflags.Model) error {
	var batteryAmount uint = 2
	banks, err := parseBanks(model.FileContent)
	if err != nil {
		return fmt.Errorf("could not parse content to banks: %w", err)
	}
	solution, err := calcSumOfBanks(banks, batteryAmount)
	if err != nil {
		return fmt.Errorf("could not calc sum of banks: %w", err)
	}
	fmt.Printf("Solution 1: %d\n", solution)
	return nil
}

func task2(model globalflags.Model) error {
	var batteryAmount uint = 12
	banks, err := parseBanks(model.FileContent)
	if err != nil {
		return fmt.Errorf("could not parse content to banks: %w", err)
	}
	solution, err := calcSumOfBanks(banks, batteryAmount)
	if err != nil {
		return fmt.Errorf("could not calc sum of banks: %w", err)
	}
	fmt.Printf("Solution 2: %d\n", solution)
	return nil
}

func parseBanks(input string) ([]Bank, error) {
	var banks []Bank
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		digits := strings.Split(line, "")
		var bank Bank
		for _, digit := range digits {
			battery, err := strconv.Atoi(digit)
			if err != nil {
				return nil, fmt.Errorf("conversion to int failed, expected digit, got %+v: %w", digit, err)
			}
			bank = append(bank, Battery(battery)) // parsing to int to uint should be safe because only single digits are parsed
		}
		banks = append(banks, bank)
	}
	return banks, nil
}

func calcSumOfBanks(b []Bank, batteryAmount uint) (uint64, error) {
	var sum uint64
	for _, bank := range b {
		joltage, err := bank.GetLargestJoltage(uint(batteryAmount))
		if err != nil {
			return 0, fmt.Errorf("can not calculate largest joltage: %w", err)
		}
		sum += joltage
	}
	return sum, nil
}
