package arrays

import (
	"sort"
)
func (a *array) MergeInterval(intervals [][]int) [][]int {

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    ans := [][]int{}

    prevStart := intervals[0][0]
    prevEnd := intervals[0][1]

    for i := 1; i < len(intervals); i++ {

        currentStart := intervals[i][0]
        currentEnd := intervals[i][1]

        if prevEnd >= currentStart {

            if currentEnd > prevEnd {
                prevEnd = currentEnd
            }

        } else {

            ans = append(ans, []int{prevStart, prevEnd})

            prevStart = currentStart
            prevEnd = currentEnd
        }
    }

    ans = append(ans, []int{prevStart, prevEnd})

    return ans

}