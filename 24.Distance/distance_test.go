package distance

import (
	"testing"
	"math"
)

func TestDist(t *testing.T){
	var x1,y1,x2,y2 float64
	x1,y1,x2,y2 = 10,10,1,1
	want := math.Sqrt(math.Pow(x1-x2,2)+math.Pow(y1-y2,2))
	first := NewPoint(x1,y1)
	second := NewPoint(x2,y2)
	got := Distance(first,second)
	if want != got{
		t.Errorf("Test failed, want %v got %v",want,got)
	}
	t.Logf("\nPassed, the distance between points is %v \n", got)
	
}