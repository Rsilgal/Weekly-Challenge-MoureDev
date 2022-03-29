package shapes

type Polygon interface {
	Area() float64
	ToString() string
}

func GetArea(p Polygon) string {
	return p.ToString()
}
