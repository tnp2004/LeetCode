package romantointeger

func romanToInt(s string) int {
	var sum int
	romanSymbol := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, c := range s {
		cur := romanSymbol[string(c)]
		if i+1 < len(s) {
			if romanSymbol[string(s[i+1])] > cur {
				sum -= cur
				continue
			}
		}
		sum += cur
	}

	return sum
}
