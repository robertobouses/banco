package informacion

import "fmt"

type DatosPersonales struct {
	dni      int
	nombre   string
	apellido string
}

var Clientes []DatosPersonales

func (d *DatosPersonales) GrabarDatosPersonales() {
	fmt.Println("Introduce los datos del Cliente, DNI, NOMBRE, APELLIDO")
	fmt.Scanln(&d.dni)
	fmt.Scanln(&d.nombre)
	fmt.Scanln(&d.apellido)
	Cliente1 := DatosPersonales{
		dni:      d.dni,
		nombre:   d.nombre,
		apellido: d.apellido,
	}

	Clientes = append(Clientes, Cliente1)
	fmt.Println("Clientes", Clientes)
}
