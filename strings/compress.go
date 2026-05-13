package stringQues

import "strconv"

func (a *stringQuestion) Compress(chars []byte) int {
	ans := []byte{}
	for i :=0; i<len(chars); i++ {
		count := 1

		for i+1 < len(chars) && chars[i] == chars[i+1]{
			count++
			i++
		}
		ans = append(ans, chars[i])

		if count > 1 {
			str := strconv.Itoa(count)
			for j := 0 ; j<len(str) ;j++ {
				ans = append(ans, str[j])
			}

		}

	}

	copy(chars,ans)

	return len(ans)

}