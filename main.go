package main

import (
	"Part2/Controller"
	"Part2/Views"
)

func main() {
	// Citim valorile de intrare folosind pachetul View
	rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx := Views.Read()

	// Determinăm soluțiile folosind controller-ul
	v, gamma, Vmax, Gammamax := Controller.DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax)

	// Afișăm rezultatele folosind View
	Views.DispRes(nil, rho, S, P, G, Cx0, k, eta, Cz, Cx, v, gamma, Vmax, Gammamax)
}
