package recursionbacktracking

func (a *rb) Partition(s string) [][]string {
	ans := [][]string{}

	findPalindrom(s, &ans, 0,[]string{})

	return ans
}	

func findPalindrom(s string, ans *[][]string, idx int, store []string){
    if len(s) == idx {
        temp := make([]string, len(store))
        copy(temp, store)
        *ans = append(*ans, temp)
        return
    }

    for i := idx ; i < len(s) ; i++ {
        if IsPalindrom(s, idx , i){
            store = append(store, s[idx:i+1])
            findPalindrom(s,ans,i+1,store)
            store = store[:len(store)-1]
        }
    }
}
func IsPalindrom(s string, left int, right int) bool {

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}