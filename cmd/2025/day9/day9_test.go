package day9

import "testing"

func TestPoint_area(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		first   Point  // Named input parameters for target function.
		another Point
		want    int64
	}{
		{
			name: "sample",
			first: Point{
				X: 2,
				Y: 5,
			},
			another: Point{
				X: 11,
				Y: 1,
			},
			want: 50,
		},
		{
			name: "line",
			first: Point{
				X: 2,
				Y: 3,
			},
			another: Point{
				X: 7,
				Y: 3,
			},
			want: 6,
		},
		{
			name: "3",
			first: Point{
				X: 7,
				Y: 1,
			},
			another: Point{
				X: 11,
				Y: 7,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.first.area(tt.another)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("area() = %v, want %v", got, tt.want)
			}
		})
	}
}
