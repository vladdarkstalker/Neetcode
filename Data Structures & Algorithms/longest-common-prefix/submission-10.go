func longestCommonPrefix(strs []string) string {
	var prefix string
	var count int
	for i := 0; i < len(strs[0]); i++ {
		var chr = strs[0][i]
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) {
				break
			}
			if chr == strs[j][i] {
				count++
			}
		}
		if count == len(strs) {
			prefix += string(chr)
			count = 0
		} else {
            break
        }
	}
	return prefix
}