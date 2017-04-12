package main

import (
	"fmt"
	"github.com/bxcodec/saint"
   "math"
)

func main() {
	 	
            
    
   arr:=[] int {2,1,3,5,6}
   var x int 

   x= saint.Max(arr...) 
   fmt.Println(x) // 6

   x= saint.Max(4,3,1,5,7) 
   fmt.Println(x) // 7

   x= saint.Min(arr...) 
   fmt.Println(x) // 1

   x= saint.Min(4,3,5,5,7) 
   fmt.Println(x) // 3

   x= saint.Sum(arr...) 
   fmt.Println(x) // 17

   x = saint.Sum(4,5)
   fmt.Println(x) // 9



   x = -5

   fmt.Println(saint.Abs(x))
   y:=float64(x)
   fmt.Println(math.Abs(y))
   x=int(y)

}


