package day1

func findSumToAndMultipy( reportArr []int , sumAmount int) int {
	if sumAmount <=0  {
		return 0
	}
	reportLen := len(reportArr)
	m  := make(map[int]bool)
	for i := 0; i < reportLen; i++ {
		searchAmount := sumAmount - reportArr[i]
		m[reportArr[i] ] = true;
		if(m[searchAmount]) {
			return reportArr[i] * searchAmount
		}
	}
	return 0

}

func find2020AndMultipy( reportArr []int) int {
	return findSumToAndMultipy(reportArr, 2020)
}

