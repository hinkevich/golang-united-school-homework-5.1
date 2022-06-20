package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	var endSqare Point
	endSqare.x = s.start.x + int(s.a)
	endSqare.y = s.start.y - int(s.a)
	return endSqare
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
