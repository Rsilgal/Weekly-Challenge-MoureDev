package shapes

import "fmt"

type Triangle struct {
	Height, Base float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) ToString() string {
	return fmt.Sprintf("Triangle.\n base: %f\n Heigth: %f\n", t.Base, t.Height)
}
