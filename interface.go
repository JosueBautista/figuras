package figuras

import (
	"fmt"
	"reflect"
)

type Figura interface {
	perimetro() float32
	area() float32
}

func CalcularPerimetro(figura Figura) {
	fmt.Printf("El perimetro del %v es: %.2f \n", reflect.TypeOf(figura).Elem().Name(), figura.perimetro())
}

func CalcularArea(figura Figura){
	fmt.Printf("El area del %v es: %.2f \n", reflect.TypeOf(figura).Elem().Name(), figura.area())
}
