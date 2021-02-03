package day7

//import (
//	"fmt"
//)

var (
	ruleMap = rules()
)

func BagCombinations(bagName string) int {
	count := 0
	for k, _ := range ruleMap {
		if k == bagName {
			continue;
		}
		count += bagDFS(k, bagName)
		//fmt.Println()
	}
	return count
}


func bagDFS(startVertex, endVertex string) int {
	visited := map[string]bool{}
	var stack []string
	stack = append(stack, startVertex)
	for len(stack) > 0 {
		//stack, sTop := pop(stack)
		n := len(stack) - 1
		sTop := stack[n]
		stack = stack[:n]
		if !visited[sTop] {
			visited[sTop] = true
			//fmt.Printf("%s ->", sTop)
			for k, _ := range ruleMap[sTop] {
				if(k == endVertex) {
					//fmt.Printf(" %s", k)
					return 1
				}
				stack = append(stack, k)
			}
		}
	}
	//fmt.Printf("//")
	return 0
}

func rules() map[string]map[string]int {
	ruleMap := make(map[string]map[string]int)

	ruleMap["light red"]    = map[string]int { "bright white" : 1, "muted yellow" : 2}
	ruleMap["dark orange"]  = map[string]int { "bright white" : 3, "muted yellow" : 4}
	ruleMap["bright white"] = map[string]int { "shiny gold" : 1}
	ruleMap["muted yellow"] = map[string]int { "shiny gold" : 2, "faded blue": 9}
	ruleMap["shiny gold"] 	= map[string]int { "dark olive" : 1, "vibrant plum": 2}
	ruleMap["dark olive"] 	= map[string]int { "faded blue" : 3, "dotted black": 4}
	ruleMap["vibrant plum"] = map[string]int { "faded blue" : 5, "dotted black": 6}
	ruleMap["faded blue"]   = map[string]int {}
	ruleMap["dotted black"] = map[string]int {}

	return ruleMap
}