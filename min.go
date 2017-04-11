package integres

// Min return the minimun  value beetwen  arg and arg0.
//
// 	Example:
//		Min(5,6) => 5
//		Min(5,5) => 5
func Min(arg int , arg0 int ) int {
	
	if arg <=arg0{
		return arg
	}
	return arg0
}