package figuras

import "math"


type Circulo struct {
	Radio float32
}

// $d PERIMETRO
func (c *Circulo) perimetro() float32 {
	return 2 * math.Pi * c.Radio
}

//* AREA
func (c *Circulo) area() float32{
	return math.Pi * c.Radio * c.Radio
}
