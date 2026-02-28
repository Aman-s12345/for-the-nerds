package main

import (
	"fmt"
	"github.com/Aman-s12345/for-the-nerds/arrays"
)

func main (){
	nums := []int{2,2,1,1,1,2,2}
	array := arrays.NewArray()
	// ans := array.MajorityElementHashMap(nums)

	ans := array.MajorityElementBoyerMoore(nums)

	
	fmt.Printf("%d",ans)


}