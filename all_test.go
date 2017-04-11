package integres

import (
	"testing"
)

func TestAbs(t *testing.T) {
	var x,y int 
	x = Abs(-15)
    y = Abs(15)
	if x != y  {
        t.Error("Expected TRUE got FALSE")    
	}
}


func TestMin(t *testing.T) {
	var x,y int 
	// z =15
	y =1
	arr:=[] int {2,1,3,5,6}
	x = Min(arr...)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", x)    
	}

	y =1
	
	x = Min(1,3)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", x)    
	}

	y =0
	
	x = Min(1,3,0)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", x)    
	}



	
}
func TestMax(t *testing.T) {
	var x,y int 
	// z =15
	y =6
	arr:=[] int {2,1,3,5,6}
	x = Max(arr...)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", x)    
	}

	y =3
	
	x = Max(1,3)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", x)    
	}

	y =3
	
	x = Max(1,3,0)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", x)    
	}


    
}

func TestPow(t *testing.T) {
	var x,y,z,result int 
	x = 3
	y = 3
	result = 27
	z = Pow(x,y)

	if z != result{
		t.Error("Expected " , result , " got " , z)
	}

}

