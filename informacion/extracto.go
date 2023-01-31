package informacion

import (
	"fmt"
	//"github.com/robertobouses/banco/movimientos"
)

type CuentaBancaria struct {
	id        int
	cliente   DatosPersonales
	gestiones []Movimiento
	saldo     int
}

type Movimiento struct {
	tipo    string
	importe int
	saldo   int
}

func Imprimir(movimientosCargo []Movimiento, movimientosAbono []Movimiento, saldo int, Clientes []DatosPersonales) {
	fmt.Println("El cliente:", Clientes)
	fmt.Println("Los movimientos de cargo son", movimientosCargo)
	fmt.Println("Los movimientos de abono son", movimientosAbono)
	fmt.Println("El saldo de la cuenta es", saldo)
}
