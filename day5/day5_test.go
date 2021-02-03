package day5

import (
	"testing"
)


func TestValidBoardingPassID(t *testing.T) {
	seatID := binaryBoardFixedSize("FBFBBFFRLR")
	if seatID != 357 {
	 t.Errorf("binaryBoardFixedSize was incorrect, got: %d, want: %d.", seatID, 357)
	}
}