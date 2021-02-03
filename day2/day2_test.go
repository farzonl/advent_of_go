package day2

import (
	"testing"
)

// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc

func TestValidPassword(t *testing.T) {
	valid := passwordPhilosophy(1, 3, 'a', "abcde")
	if !valid {
	 t.Errorf("passwordPhilosophy was incorrect, got: %t, want: %t.", valid, true)
	}
}

func TestInvalidPassword(t *testing.T) {
	valid := passwordPhilosophy(1, 3, 'b', "cdefg")
	if valid {
	 t.Errorf("passwordPhilosophy was incorrect, got: %t, want: %t.", valid, false)
	}
}

func TestValidPasswordDuplicates(t *testing.T) {
	valid := passwordPhilosophy(2, 9, 'c', "ccccccccc")
	if !valid {
	 t.Errorf("passwordPhilosophy was incorrect, got: %t, want: %t.", valid, true)
	}
}
