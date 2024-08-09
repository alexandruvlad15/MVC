package Controller

import (
	"Part2/Model"
)

func DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax float64) (v, gamma, Vmax, Gammamax float64) {
	solutions := Model.Optimize(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax)
	v = solutions[0]
	gamma = solutions[1]
	Vmax = Model.CalcVmax(G, rho, S, Czmax)
	Gammamax = Model.CalcGammamax(P, eta, rho, S, Vmax, Cx0, k, Czmax, G)
	return
}
