package movimientos

import "fmt"

func LeerMonto() int {
	var monto int
	fmt.Println("Indica el monto de la operación")
	fmt.Scanln(&monto)
	return monto
}
