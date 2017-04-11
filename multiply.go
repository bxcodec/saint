package intregers

// Returns the multiplication result from the  integer's  argument,  between arg to arg0.
//
// Example : 
//  	x:=Multiply(2,3) // 6
//  	x:=Multiply(3,3) // 9
func Multiply(arg int , arg0 int) int {
	
	res := arg * arg0
	
	if res <0 && ((arg >0 && arg0 >0) || (arg <0 && arg0 <0) ){
		panic("Integer overflow, use 'math' library for more big calculations")
	}else
	if res > 0 && ((arg >0 && arg0 <0) || (arg <0 && arg0 >0) ){
		panic("Integer overflow, use 'math' library for more big calculations")
	}

	return res
}
