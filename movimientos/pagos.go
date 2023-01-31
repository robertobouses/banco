package movimientos

import "fmt"

func (c *CuentaBancaria) Abono(monto int) ([]Movimiento, int) {
	c.saldo -= monto
	c.gestiones = append(c.gestiones, Movimiento{tipo: "ABONO", importe: monto, saldo: c.saldo})
	fmt.Println("El movimiento es del tipo CARGO por un importe de:", monto)
	fmt.Println("El saldo de la cuenta es:", c.saldo)
	return c.gestiones, c.saldo
}
