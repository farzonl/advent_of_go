package day1

import "testing"

func TestFindAndMultipy(t *testing.T) {
	report := []int {1721, 979, 366, 299, 675, 1456}
	total := find2020AndMultipy(report)
	if total != 514579 {
	 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 514579)
	}
}

func TestFindAndMultipyEmptyArr(t *testing.T) {
	report := []int {}
	total := find2020AndMultipy(report)
	if total != 0 {
	 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestFindAndMultipyNegative(t *testing.T) {
	report := []int {1,2,3}
	total := findSumToAndMultipy(report,-1)
	if total != 0 {
	 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestFindAndMultipyZero(t *testing.T) {
	report := []int {1,2,3}
	total := findSumToAndMultipy(report,0)
	if total != 0 {
	 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestFindAndMultipyOneSumZero(t *testing.T) {
	report := []int {1,2,3,0}
	total := findSumToAndMultipy(report,1)
	if total != 0 {
	 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestFindAndMultipyOneNoSum(t *testing.T) {
	report := []int {1,2,3}
	total := findSumToAndMultipy(report,1)
	if total != 0 {
	 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 0)
	}
}