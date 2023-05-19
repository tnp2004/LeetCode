package findSmallestLetterGreaterThanTarget

func nextGreatestLetter(letters []byte, target byte) byte {
	if target == 122 {
		return letters[0]
	}

	for _, cb := range letters {
		if cb > target {
			return cb
		}
	}

	return letters[0]
}
