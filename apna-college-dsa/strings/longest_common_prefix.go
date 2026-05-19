package stringQues


func (a *stringQuestion) LongestCommonPrefix(strs []string) string {
    
	if len(strs) == 0 {
		return ""
	}
	count := 0

	for i:=0 ; i < len(strs[0]) ; i++ {
		char := strs[0][i]

		for j := 1 ; j < len(strs) ; j++ {
			if i >= len(strs[j]) || char != strs[j][i] {
				return strs[0][:count]
			}else {
				count++
			}
		}
	}
	return strs[0]
   
}