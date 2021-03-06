package saint



// Returns the power  value of integer's  argument from arg to arg0.
//
// Example : 
//  	x:=Pow(2,3) // 8
//  	x:=Pow(3,3) // 27
func Pow(arg int , arg0 int) int {
	if arg0 < 0 {
		return 0
	}
	res := 1
	for arg0 != 0 {
		if (arg0 & 1) == 1 {
			res *= arg
		}
		arg0 >>= 1
		arg *= arg
	}

	if res <0{
		panic("Integer overflow, use 'math' library for more big calculations")
	}
	return res
}
