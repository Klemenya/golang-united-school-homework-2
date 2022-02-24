package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type myInt int

const SidesTriangle myInt = 3
const SidesSquare myInt  = 4
const SidesCircle myInt = 0




func CalcSquare(sideLen float64, sideNum myInt) float64 {
	switch sideNum {
	case SidesSquare:
		return (sideLen * sideLen)
	case SidesTriangle:
		return ((math.Sqrt(3) / 4) * (sideLen * sideLen))
	case SidesCircle:
		return (math.Pi * (sideLen * sideLen))
	default:
		return 0
	}
}

