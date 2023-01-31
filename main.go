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
	var monto, saldo int
	var movimientosAbono []movimientos.Movimiento
	var movimientosCargo []movimientos.Movimiento

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
			fmt.Println("VA A INTRODUCIR UN INGRESO")
			monto = movimientos.LeerMonto()
			movimientosCargo, saldo = valor.Cargo(monto)

		case 2:
			fmt.Println("VA A INTRODUCIR UN GASTO")
			monto = movimientos.LeerMonto()
			movimientosAbono, saldo = valor.Abono(monto)

		case 3:
			fmt.Println("OBTENIENDO EXTRACTO")
			informacion.Imprimir(movimientosCargo, movimientosAbono, saldo, informacion.Clientes)

		case 4:

			fmt.Println("Va a grabar los datos personales")
			informacion.Clientes = info.GrabarDatosPersonales()

		case 5:
			return

		default:
			fmt.Println("No está en el menú")
		}
	}
}
