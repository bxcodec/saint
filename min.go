package integres

// Min return the minimun  value beetwen  arg and arg0.
//
// 	Example:
//		x:=Min(5,6) //  x = 5
//		y:=Min(5,5) //  y = 5
func Min(arg int , arg0 int ) int {
	
	if arg <=arg0{
		return arg
	}
	return arg0
}