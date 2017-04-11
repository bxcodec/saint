package intregers



// Return the minimun  value beetwen  all the integer's arguments
//
// 	Example:
//		x:=Min(5,6,3,1,2,6) //  x = 1
//		y:=Min(5,5,32,2) //  y = 2
//
//		arr:=[] int {2,1,3,5,6}
//		x = Min(arr...) // x = 1
func Min(arg ...int  ) int {
	
	res := findMin(arg)	

	return res
}

func findMin(a [] int)  int{
	var min int = a[0] 
	for _,arg := range a{
		if (min>arg){
			min = arg
		}
	}
	return min
}

