package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

type Side int

func CalcSquare(sideLen float64, sidesNum Side) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * (sideLen * sideLen)
	case SidesSquare:
		return sideLen * sideLen
	case SidesTriangle:
		s := (sideLen * SidesTriangle) / 2
		return math.Sqrt(s * (s - sideLen) * (s - sideLen) * (s - sideLen))
	default:
		return 0
	}
}
