package stringQues


type stringQuestion struct{}

func NewString() String {
	return &stringQuestion{}
}

func (a *stringQuestion)  IsPalindrome(s string) bool {
	left := 0 
    right := len(s)-1

    for left < right {
        // Skip non alphanumeric chars
        for left < right && !isAplhaNumeric(s[left]){
            left++
        }

        for right > left && !isAplhaNumeric(s[right]){
            right--
        }

        if convertLower(s[left]) != convertLower(s[right]){
            return false
        }
        left++
        right--


    }
    return true

}

func isAplhaNumeric(ch byte) bool {
    // small letter
    if ch >= 'a' && ch <= 'z' {
        return true
    }
    // capital letter
    if ch >= 'A' && ch <= 'Z' {
        return true
    }
    // numbers
    if ch >=  '0' && ch <= '9' {
        return true
    }
    return false
}

func convertLower(ch byte) byte {
    if ch >= 'A' && ch <='Z' {
        ch = ch + ('a' - 'A')
    }
    return ch
}