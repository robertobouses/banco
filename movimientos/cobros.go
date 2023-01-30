package movimientos

import "github.com/robertobouses/banco/informacion"

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

func (c *CuentaBancaria) Cargo(monto int) {
	c.saldo += monto

}
