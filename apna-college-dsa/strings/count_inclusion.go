package stringQues

func (a *stringQuestion) CheckInclusionWithPermutation(s1 string, s2 string) bool {
	permute := make(map[string]bool)
	chars := []byte(s1)

	generatePermutation(chars, 0, permute)
	windowSize := len(s1)

	for i := 0; i <= len(s2)-windowSize; i++ {
		sub := s2[i : i+windowSize]

		if permute[sub] {
			return true
		}
	}

	return false

}

func generatePermutation(chars []byte, idx int, permute map[string]bool) {
	if idx == len(chars) {
		permute[string(chars)] = true
		return
	}

	for i := idx; i < len(chars); i++ {
		// swap
		chars[idx], chars[i] = chars[i], chars[idx]
		// generate new permute
		generatePermutation(chars, idx+1, permute)
		//swap back to main bytes
		chars[idx], chars[i] = chars[i], chars[idx]
	}
}

func (a *stringQuestion) CheckInclusionOptmized(s1 string, s2 string) bool {
	count1 := [26]int{}
	count2 := [26]int{}

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i<len(s1) ; i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for i := len(s1) ; i < len(s2) ; i++ {
		count2[s2[i]-'a']++
		count2[s2[i-len(s1)]-'a']--
		if count1 == count2 {
			return true
		}
	}

	return false


}