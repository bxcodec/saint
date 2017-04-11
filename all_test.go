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
	var x,y,z int 
	z =15
	y =-15
	x = Min(z,y)
    
	if x != y  {
        t.Error("Expected ", y ,  "got ", z)    
	}

	z =-15
	y =15
	x = Min(z,y)

	if x != z  {
        t.Error("Expected ", z ,  "got ", y)    
	}

    
    z =-15
	y =-15
	x = Min(z,y)

	if x != z || x != y  {
        t.Error("Expected ", y ,  "got ", z)    
	}
    
}

