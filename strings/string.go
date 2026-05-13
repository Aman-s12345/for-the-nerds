package stringQues

type String interface {
	// valid palindrom
	IsPalindrome(s string) bool

	// valid anagram
	IsAnagramWithMap(s string, t string) bool 
	IsAnagramWithArray(s string, t string) bool 

	// reverse Words
	ReverseWords(s string) string 

	// remove Occurrance
	RemoveOccurrences(s string, part string) string 

	// 


}