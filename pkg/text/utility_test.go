package text

import (
	"fmt"
	"testing"
)

func TestLamportsToSol(t *testing.T) {
	var tests = []struct {
		have uint64
		want uint64
	}{
		{523044252942, 523},
		{500000000000, 500},
		{580000000000, 580},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.have, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := LamportsToSol(tt.have)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSolToLamports(t *testing.T) {
	var tests = []struct {
		have uint64
		want uint64
	}{
		{523, 523000000000},
		{500, 500000000000},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.have, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := SolToLamports(tt.have)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
