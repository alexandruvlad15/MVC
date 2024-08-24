package Controller1

import (
	"Science/Model1"
)

func Det_portanta(rho float64, v float64, Cz float64) float64 {
	var P float64
	p1, p2, p3 := Model1.Det_pol()
	S := Model1.Det_suprafata(p1, p2, p3) //determinarea suprafetei
	P = Model1.Det_P(rho, v, S, Cz)       //determinarea fortei de portanta
	return P
}

func Det_rezistenta(rho float64, v float64, Cz float64) float64 {
	var D float64
	p1, p2, p3 := Model1.Det_pol()
	S := Model1.Det_suprafata(p1, p2, p3) //determinarea suprafetei
	D = Model1.Det_D(rho, v, S, Cz)       //determinarea fortei de portanta
	return D
}
