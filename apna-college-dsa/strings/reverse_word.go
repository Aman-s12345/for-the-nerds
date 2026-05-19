package stringQues

func (a *stringQuestion) ReverseWords(s string) string {

	words := []string{}

	i := 0
	for i < len(s) {

		for i < len(s) && s[i] == ' ' {
			i++
		}
		start := i

		for i < len(s) && s[i] != ' ' {
			i++
		}

		if start < i {
			words = append(words, s[start:i])
		}

	}

	ans := ""

	for i := len(words) - 1; i >= 0; i-- {
		ans += words[i]

		if i != 0 {
			ans += " "
		}
	}

	return ans
}
