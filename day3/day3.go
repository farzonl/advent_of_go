
package day3

import (
	"bufio"
	//"fmt"
	"os"
)


func parseInput(path string) []string {
	//path := "C:\\Users\\Farzon\\Projects\\advent_of_go\\day3\\input.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	
	var parsedLines []string
	fScan := bufio.NewScanner(file)
	fScan.Split(bufio.ScanLines) // split on newlines
	for fScan.Scan() {
		parsedLines = append(parsedLines, fScan.Text())
	}
	file.Close()

	return parsedLines
}

func countTrees( parsedLines []string, xDir, yDir int) int {

	treeCount := 0
	x := 0
	for i, line := range parsedLines {
		if(i % yDir != 0) {
			continue
		}
		//if(i % yDir == 0 && line[x % len(line)] == '#') {
		if(line[x % len(line)] == '#') {
			treeCount++
		}
		x += xDir
	}
	return treeCount
}

