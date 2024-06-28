package calc_distance

import (
	"testing"
)

type testCase struct {
	name string
	in string
	want int
}

type testCalcCase struct {
	name string
	in []string
	want int
}

func TestParseDistance(t *testing.T) {
	tests := []testCase {
		{"1", "100m", 100},
		{"2", "2km", 2000},
		{"3", "straight 0.4km", 400},
		{"4", "100m forward", 100},
		{"5", "2.4km backward", 2400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseDistance(tt.in)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}

func TestCalcDistance(t *testing.T) {
	tests := []testCalcCase {
		{"1", 
		[]string{"100m to intersection", "turn right", "straight 300m", "enter motorway", "straight 5km", 
		"exit motorway", "500m straight", "turn sharp left", "continue 100m to destination"}, 
		6000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcDistance(tt.in)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
