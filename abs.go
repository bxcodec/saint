// This library only provide some of function that "math" library doesn't support especially in <b>integer</b>'s operation.
// Intregres also provide some function to ease in some of operations. 
package intregers




// Returns the absolute value of  integer's value from argument.
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
