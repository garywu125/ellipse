package ellipse

import (
	"fmt"
	"math"

	"github.com/garywu125/ellipse/hello"
)

// if center is 0,0
type Init struct {
	A, B float64
}

// Get Eccentricity of Ellipse
func (e *Init) GetEccentricity() float64 {
	fmt.Println("version v0.1.2")
	return (math.Sqrt(math.Pow(e.A, 2) - math.Pow(e.B, 2))) / e.A
}

func Sayhello() {
	hello.Greeting()
}
