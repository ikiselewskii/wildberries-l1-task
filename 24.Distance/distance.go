package distance

import "math"

type Point struct{
	x float64
	y float64
}


func NewPoint(x,y float64)*Point{
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

func Distance(p1,p2 *Point)float64{
	d := math.Sqrt(math.Pow(p1.x-p2.x,2)+math.Pow(p1.y-p2.y,2))
	return d
}