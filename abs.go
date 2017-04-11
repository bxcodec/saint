package integres

// Returns the absolute value of  integer value from argument.
//
// 	Example:
//		x:=Abs(-5) // x = 5
//		x:=Abs(5)  // x = 5
func Abs( arg int ) int {
	
	if arg < 0{
		arg = -arg
	}
	
	return arg 
}
