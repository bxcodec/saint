package integres

// Max return the maximun  value beetwen  arg and arg0.
//
// 	Example:
//		x:=Min(5,6,3,1,2,6) //  x = 1
//		y:=Min(5,5,32,2) //  y = 2
//		arr:=[] int {2,1,3,5,6}
//		x = Min(arr...) // x = 1
func Max(arg ...int  ) int {
	
	res := findMax(arg)	

	return res
}

func findMax(a [] int)  int{
	var max int = 0
	for _,arg := range a{
		if (max<arg){
			max = arg
		}
	}
	return max
}
