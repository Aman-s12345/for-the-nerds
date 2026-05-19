package stringQues

import  (
	"sort"
)

type pair struct {
	idx int
	str string
}

func (a *stringQuestion) GroupAnagramsWithTravesal(strs []string) [][]string {
	pairs := []pair{}
	for i := 0; i < len(strs); i++ {
		byteStr := []byte(strs[i])
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})

		pairs = append(pairs, pair{
			idx: i,
			str: string(byteStr),
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].str < pairs[j].str
	})

	ans := [][]string{}
	for i := 0; i < len(pairs); i++ {
		temp := []string{}
		temp = append(temp, strs[pairs[i].idx])

		j := i + 1
		for j < len(pairs) && pairs[i].str == pairs[j].str {
			temp = append(temp, strs[pairs[j].idx])
			j++
		}
		ans = append(ans, temp)
		j = i

	}
	return ans

}


func (a *stringQuestion) GroupAnagramsWithMap(strs []string) [][]string {
	lookup := map[string][]string{}

	for i :=0; i<len(strs) ; i++ {
		byteStr := []byte(strs[i])
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})
		key := string(byteStr)
		lookup[key] = append(lookup[key], strs[i])

	}

	ans := [][]string{}

	for _,val := range lookup {
		ans = append(ans, val)
	}
	return ans

}