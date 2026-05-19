package arrays

func (a *array) MyPowNaive(x float64, n int) float64 {
	var ans = 1.0

	convert := n

	if x == 1 {
		return x
	}
	if convert < 0 {
		convert = -1 * convert
	}
	for i := 0; i < convert; i++ {
		ans = ans * x
	}
	if n > 0 {
		return ans
	} else if n < 0 {
		return 1 / ans
	}

	return 1
}


func(a *array) MyPowOptmized(x float64, n int) float64 {
    if x == 1.0 {
        return x
    }
    if n == 0 {
        return 1.0
    }
   convert := n
   if convert < 0 {
        convert = convert * -1
   }
   ans := findPow(x,convert )

   if n >0 {
    return ans
   }else {
    return 1.0/ans
   }
}

func findPow(x float64, n int) float64 {
    if (n == 0){
        return 1
    }
    if (n == 1){
        return x
    }
    
    helper := findPow(x, n/2)
    helper *= helper
    if n %2 == 0 {
        return helper
    }else {
        return x*helper
    }
}
