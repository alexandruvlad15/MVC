package main

import (
	"Part2/Controller" //pentru aflarea rezultatelor
	"Part2/Views"      //pentru citirea si afisarea datelor
)

func main() {
	// input
	rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx := Views.Read()

	// obtinerea rezultatelor
	v, gamma, Vmax, Gammamax := Controller.DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax)

	// output
	Views.DispRes(nil, rho, S, P, G, Cx0, k, eta, Cz, Cx, v, gamma, Vmax, Gammamax)
}
