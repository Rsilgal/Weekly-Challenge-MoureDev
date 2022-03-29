package shapes

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) ToString() string {
	return fmt.Sprintf("Rectangle.\n Width: %f\n Heigth: %f\n", r.Width, r.Height)
}
