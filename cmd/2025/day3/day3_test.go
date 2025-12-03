package day3_test

import (
	"testing"

	"github.com/marceljk/advent-of-code/cmd/2025/day3"
)

func TestBank_GetLargestJoltage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank           day3.Bank
		amoutBatteries uint
		want           uint64
		wantErr        bool
	}{
		{
			name:           "example - amount 2",
			bank:           day3.Bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			amoutBatteries: 2,
			want:           98,
		},
		{
			name:           "example - amount 5",
			bank:           day3.Bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			amoutBatteries: 5,
			want:           98765,
		},
		{
			name:           "example - amount 12",
			bank:           day3.Bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			amoutBatteries: 12,
			want:           987654321111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.bank.GetLargestJoltage(tt.amoutBatteries)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetLargestJoltage() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetLargestJoltage() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetLargestJoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}
