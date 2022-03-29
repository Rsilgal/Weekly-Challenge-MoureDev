package shapes

import "fmt"

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) ToString() string {
	return fmt.Sprintf("Square.\n Side: %f\n", s.Side)
}
