package stringQues

type String interface {
	// valid palindrom
	IsPalindrome(s string) bool

	// valid anagram
	IsAnagramWithMap(s string, t string) bool 
	

}