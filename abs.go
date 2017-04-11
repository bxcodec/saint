package integres

// Abs returns the absolute value of x.
//
// 	Special cases are:
//		Abs(±Inf) => Inf
func Abs( arg int ) int {
	
	if arg < 0{
		arg = -arg
	}
	
	return arg 
}
