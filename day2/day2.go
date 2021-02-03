package day2

func passwordPhilosophy(min, max int, char byte, password string) bool {
	m  := make(map[byte]int)

	passwordLen := len(password)
	for i := 0; i < passwordLen; i++ {
		m[password[i]]++
	}
	if(m[char] < min || m[char] > max) {
		return false
	}
	return true
}