package stringQues

func (a *stringQuestion) RemoveOccurrences(s string, part string) string {
	ans := []byte{} // treat it like a stack
	for i := 0 ; i < len(s) ; i++ {
		// push in the stack
		ans = append(ans, s[i])

		// check if we found a substring
		match := true
		j := 0
		for len(ans) > 0 &&  j < len(part) {
			if (ans[len(ans)-len(part)+i] != part[j]){
				match = false
				break
			}
			j++

		}

		// remove the substring
		if match {
			ans = ans[:len(ans)-len(part)]
		}
	}
	return string(ans)

}