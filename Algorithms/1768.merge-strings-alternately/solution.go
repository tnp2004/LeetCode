package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	ls := l1 + l2
	var ms string

	for i := 0; i < ls; i++ {
		if i < l1 {
			ms += string(word1[i])
		}
		if i < l2 {
			ms += string(word2[i])
		}
	}
	return ms
}
