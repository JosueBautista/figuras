package figuras

type Cuadrado struct {
	Lado float32
}

// $d PERIMETRO
func (c *Cuadrado) perimetro() float32{
	return c.Lado * 4
}

//* AREA
func (c *Cuadrado) area() float32 {
	return c.Lado * c.Lado
}

