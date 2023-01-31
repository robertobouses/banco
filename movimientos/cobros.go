package movimientos

import (
	"fmt"

	"github.com/robertobouses/banco/informacion"
)

type CuentaBancaria struct {
	id        int
	cliente   informacion.DatosPersonales
	gestiones []Movimiento
	saldo     int
}

type Movimiento struct {
	tipo    string
	importe int
	saldo   int
}

/*
type DatosPersonales struct {
	dni      int
	nombre   string
	apellido string
}*/

func (c *CuentaBancaria) Cargo(monto int) ([]Movimiento, int) {
	c.saldo += monto
	c.gestiones = append(c.gestiones, Movimiento{tipo: "CARGO", importe: monto, saldo: c.saldo})
	fmt.Println("El movimiento es del tipo CARGO por un importe de:", monto)
	fmt.Println("El saldo de la cuenta es:", c.saldo)
	return c.gestiones, c.saldo
}
