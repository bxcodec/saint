package integres

// Return the maximun  value beetwen  all the integers arguments
//
// 	Example:
//		x:=Max(5,6,3,1,2,6) //  x = 6
//		y:=Max(5,5,32,2) //  y = 32
//
//		arr:=[] int {2,1,3,5,6}
//		x = Max(arr...) // x = 6
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
