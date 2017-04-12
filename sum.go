package saint




// Returned sum of integers  from the arguments.
// Sum  array's value from the arguments.
// Sum between two or more arguments.
//
// 	Example:
//		x:=Sum(5,6,3,1,2,6) //  x = 23
//		y:=Sum(5,2) //  y = 7
//
//		arr:=[] int {2,1,3,5,6}
//		x = Sum(arr...) // x =17
func Sum(arg ...int  ) int {
	
	res := sumValue(arg)	
	return res
}

func sumValue(a [] int)  int{
	var total int =0
	for _,arg := range a{
		total+=arg
	}
	return total
}

