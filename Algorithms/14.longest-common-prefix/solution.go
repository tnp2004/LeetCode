package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	var s string

	for i := 0; i < len(strs[0]); i++ {
		for _, v2 := range strs {
			if i == len(v2) || v2[i] != strs[0][i] {
				return s
			}
		}
		s += string(strs[0][i])
	}

	return s
}
