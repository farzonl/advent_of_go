
package day6

func parseForms(forms string) []map[byte]bool {
	formAnsSets := []map[byte]bool{}
	currSet := map[byte]bool {}
	formsLen := len(forms)
	for i := 0; i < formsLen; i++ {
		if forms[i] == '\n' {
			if i < formsLen -2 && forms[i + 1] == ' ' && forms[i + 2] == '\n' {
				formAnsSets = append(formAnsSets, currSet)
				currSet = map[byte]bool {}
				i+=2 // skip the space
			}
			continue;
		}
		currSet[forms[i]] = true
	}
	formAnsSets = append(formAnsSets, currSet)
	return formAnsSets
}

func sumOfCounts(forms string) int {
	formAnsSets := parseForms(forms)
	formsLen := len(formAnsSets)
	count := 0
	for i := 0; i < formsLen; i++ {
		count += len(formAnsSets[i])
	}

	return count
}

func sumOfCountsV2(forms string) int {
	currSet := map[byte]bool {}
	formsLen := len(forms)
	count := 0
	for i := 0; i < formsLen; i++ {
		if forms[i] == '\n' {
			if i < formsLen -2 && forms[i + 1] == ' ' && forms[i + 2] == '\n' {
				count += len(currSet)
				currSet = map[byte]bool {}
				i+=2 // skip the space
			}
			continue;
		}
		currSet[forms[i]] = true
	}
	count += len(currSet)
	return count
}