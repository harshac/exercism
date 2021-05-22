// Package triangle specifies different kinds of triangles
package triangle

import "math"

//Kind specifies the type of triangle
type Kind string

const (
	//NaT kind represents not a triangle
	NaT = "NaT"
	//Equ kind represents equilateral triangle
	Equ = "Equ"
	//Iso kind represents isoceles triangle
	Iso = "Iso"
	//Sca kind represents scalene triangle
	Sca = "Sca"
)

// KindFromSides determines the kind of triangle
func KindFromSides(a, b, c float64) Kind {
	var isValid = func(n float64) bool {
		return !(n == 0 || math.IsNaN(n) || math.IsInf(n, 1) || math.IsInf(n, -1))
	}
	if !(isValid(a) && isValid(b) && isValid(c)) {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
