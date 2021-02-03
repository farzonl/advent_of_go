package day3

import (
	"testing"
)

func TestValidTreeCountBasic(t *testing.T) {
	parsedLines := parseInput("basic.txt")
	count := countTrees(parsedLines,3,1)
	ans :=7
	if count != ans {
		t.Errorf("countTrees was incorrect, got: %d, want: %d.", count, ans)
	}
}

func TestValidTreeCountInput1(t *testing.T) {
	parsedLines := parseInput("input1.txt")
	ans := []int{77,280,74,78,35}
	boundsArr := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for i,  b := range boundsArr {
		count := countTrees(parsedLines,b[0],b[1])
		if count != ans[i] {
			t.Errorf("countTrees was incorrect, got: %d, want: %d.", count, ans[i])
		}
	}
}

func TestValidTreeCountInput2(t *testing.T) {
	parsedLines := parseInput("input2.txt")
	ans := []int{65,237,59,61,38}
	boundsArr := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for i,  b := range boundsArr {
		count := countTrees(parsedLines,b[0],b[1])
		if count != ans[i] {
			t.Errorf("countTrees was incorrect, got: %d, want: %d.", count, ans[i])
		}
	}
}

func TestValidTreeCountInput3(t *testing.T) {
	parsedLines := parseInput("input3.txt")
	ans := []int{66,200,76,81,46}
	boundsArr := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for i,  b := range boundsArr {
		count := countTrees(parsedLines,b[0],b[1])
		if count != ans[i] {
			t.Errorf("countTrees was incorrect, got: %d, want: %d.", count, ans[i])
		}
	}
}

func TestValidTreeCountInput4(t *testing.T) {
	parsedLines := parseInput("input4.txt")
	
	ans := []int{84,195,70,70,47}
	boundsArr := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for i,  b := range boundsArr {
		count := countTrees(parsedLines,b[0],b[1])
		if count != ans[i] {
			t.Errorf("countTrees was incorrect, got: %d, want: %d.", count, ans[i])
		}
	}
}
