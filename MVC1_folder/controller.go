package Controller

import (
	"Part2/Model" //pentru a utiliza algoritmul de determinare a parametrilor
)

func DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax float64) (v, gamma, Vmax, Gammamax float64) { //functie ce determina valorile celor 4 parametri
	solutions := Model.Optimize(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax) //functia de optimizare, ce determina vectorul de solutii pentru v si gamma
	v = solutions[0]
	gamma = solutions[1]
	Vmax = Model.CalcVmax(G, rho, S, Czmax)
	Gammamax = Model.CalcGammamax(P, eta, rho, S, Vmax, Cx0, k, Czmax, G)
	return
}
