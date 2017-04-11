package integres

// Abs returns the absolute value of arg.
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
