package arrays


func (a *array)LengthOfLongestSubstring(s string) int {
	lookUP := map[byte]int{}
	maxAns := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		char := s[windowEnd]

		if lastSeen, ok := lookUP[char]; ok && lastSeen >= windowStart {
			windowStart = lastSeen + 1
		}

		currentAns := windowEnd - windowStart + 1
		maxAns = max(maxAns, currentAns)

		lookUP[char] = windowEnd
	}

	return maxAns
}