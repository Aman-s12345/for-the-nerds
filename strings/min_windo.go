package stringQues

import "math"

func (a *stringQuestion) MinWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(need)

	left := 0
	start := 0
	formed := 0

	window := make(map[byte]int)

	minLen := math.MaxInt32

	for right := 0; right < len(s); right++ {

		// expand window
		char := s[right]
		window[char]++

		// check if requirement satisfied
		if need[char] > 0 && window[char] == need[char] {
			formed++
		}

		// shrink window
		for formed == required {

			// update answer
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			leftChar := s[left]
			window[leftChar]--

			// check if window became invalid
			if need[leftChar] > 0 &&
				window[leftChar] < need[leftChar] {
				formed--
			}

			left++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[start : start+minLen]
}