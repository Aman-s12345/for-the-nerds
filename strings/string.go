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

	// check Inclusion
	CheckInclusionWithPermutation(s1 string, s2 string) bool

	// Compress
	Compress(chars []byte) int
	// longest Common Prefix
	LongestCommonPrefix(strs []string) string

	// Group Anagrams
	GroupAnagramsWithTravesal(strs []string) [][]string 

}

