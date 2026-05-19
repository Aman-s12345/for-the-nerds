package stringQues

func (a *stringQuestion) IsAnagramWithMap(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	lookup := map[byte]int{}

	for i := 0; i < len(s); i++ {
		lookup[s[i]]++
	}

	for i := 0; i < len(t); i++ {

		if lookup[t[i]] == 0 {
			return false
		}

		lookup[t[i]]--
	}

	return true
}

func (a *stringQuestion) IsAnagramWithArray(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
