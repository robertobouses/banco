package main

import (
	"github.com/robertobouses/banco/informacion"

	"github.com/robertobouses/banco/movimientos"
)

func main() {

	movimientos.Cargo()
	informacion.GrabarDatosPersonales()
}
