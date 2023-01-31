package movimientos

func (c *CuentaBancaria) Abono(monto int) {
	c.saldo -= monto
}
