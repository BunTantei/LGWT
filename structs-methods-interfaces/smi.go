// Package structs-methods-interfaces to learn about structs, methods, and interfaces.
package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Shape interface - any type with an Area() method is a Shape
type Shape interface {
	Area() float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

// Area method on Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method on Circle
func (c Circle) Area() float64 {
	return 3.141592653589793 * c.Radius * c.Radius
}

// Area method on Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
