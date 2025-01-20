package exercises

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int32

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// check if it is a triangle
	if !(a+b >= c && b+c >= a && a+c >= b) || a == 0 || b == 0 || c == 0 {
		return NaT
	}

	if a == b && b == c && a == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}
