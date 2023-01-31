package main

import (
	"fmt"

	"github.com/robertobouses/banco/informacion"

	"github.com/robertobouses/banco/movimientos"
)

func main() {

	valor := movimientos.CuentaBancaria{}
	info := informacion.DatosPersonales{}

	var opcion int
	var monto int

	for {

		fmt.Println("1. Cargar")
		fmt.Println("2. Abonar")
		fmt.Println("3. Extracto")
		fmt.Println("4. Grabar Datos Personales")
		fmt.Println("5. Salir")

		fmt.Print("Selecciona una opción: ")

		fmt.Scanln(&opcion)
		switch opcion {

		case 1:

			fmt.Println("indique el monto")
			fmt.Scanln(&monto)
			valor.Cargo(monto)

		case 2:
			fmt.Println("indique el monto")
			fmt.Scanln(&monto)
			valor.Abono(monto)

		case 3:
			informacion.Imprimir()

		case 4:

			fmt.Println("Va a grabar los datos personales")
			info.GrabarDatosPersonales()

		case 5:
			return

		default:
			fmt.Println("No está en el menú")
		}
	}
}
