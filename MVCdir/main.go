package main

import (
	"Science/Controller1"
	"Science/View1"
)

func main() {
	rho, v, Cz, Cx := View1.Read() //citire
	var L, D float64
	L = Controller1.Det_portanta(rho, v, Cz)   //portanta
	D = Controller1.Det_rezistenta(rho, v, Cx) //rezistenta la inaintare
	View1.DispRes(rho, v, Cz, Cx, L, D)        //afisare
}
