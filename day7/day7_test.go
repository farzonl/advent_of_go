package day7

import (
	"testing"
)

func TestValid(t *testing.T) {
	// dark orange ->muted yellow -> shiny gold
	// muted yellow -> shiny gold
	// light red ->muted yellow -> shiny gold
	// bright white -> shiny gold
	count := BagCombinations("shiny gold")
	if count != 4 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 4)
	}

	// light red ->muted yellow ->faded blue ->shiny gold -> dark olive    
	// bright white ->shiny gold -> dark olive
	// dark orange ->muted yellow ->faded blue ->shiny gold -> dark olive  
	// muted yellow ->faded blue ->shiny gold -> dark olive
	// shiny gold -> dark olive
	count = BagCombinations("dark olive")
	if count != 5 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 5)
	}

	// light red ->muted yellow ->faded blue ->shiny gold -> vibrant plum  
	// bright white ->shiny gold -> vibrant plum
	// dark orange ->muted yellow ->faded blue ->shiny gold -> vibrant plum
	// muted yellow ->faded blue ->shiny gold -> vibrant plum
	// shiny gold -> vibrant plum
	count = BagCombinations("vibrant plum")
	if count != 5 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 5)
	}

	// light red -> muted yellow
	// dark orange -> muted yellow
	count = BagCombinations("bright white")
	if count != 2 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 2)
	}
	// light red -> muted yellow
	// dark orange -> muted yellow
	count = BagCombinations("muted yellow")
	if count != 2 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 2)
	}
	
	// bright white ->shiny gold ->vibrant plum -> faded blue
	// dark olive -> faded blue
	// light red ->muted yellow -> faded blue
	// muted yellow -> faded blue
	// shiny gold ->vibrant plum -> faded blue
	// vibrant plum -> faded blue
	// dark orange ->bright white ->shiny gold ->vibrant plum -> faded blue
	count = BagCombinations("faded blue")
	if count != 7 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 7)
	}

	//dark orange ->muted yellow ->shiny gold ->dark olive -> dotted black
	//muted yellow ->shiny gold ->vibrant plum -> dotted black
	//shiny gold ->vibrant plum -> dotted black
	//vibrant plum -> dotted black
	//light red ->muted yellow ->faded blue ->shiny gold ->vibrant plum -> dotted black
	//bright white ->shiny gold ->vibrant plum -> dotted black
	//dark olive -> dotted black
	count = BagCombinations("dotted black")
	if count != 7 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 7)
	}

	count = BagCombinations("light red")
	if count != 0 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 0)
	}

	count = BagCombinations("dark orange")
	if count != 0 {
		t.Errorf("BagCombinations was incorrect, got: %d, want: %d.", count, 0)
	}
}

