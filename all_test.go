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
