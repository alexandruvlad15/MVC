package Views

import (
	"fmt" //pentru citire si afisare
)

func Read() (rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx float64) { //pentru citirea datelor(parametrilor)
	fmt.Print("rho=")
	fmt.Scan(&rho)
	fmt.Print("G=")
	fmt.Scan(&G)
	fmt.Print("S=")
	fmt.Scan(&S)
	fmt.Print("P=")
	fmt.Scan(&P)
	fmt.Print("Cx0=")
	fmt.Scan(&Cx0)
	fmt.Print("eta=")
	fmt.Scan(&eta)
	fmt.Print("k=")
	fmt.Scan(&k)
	fmt.Print("Cz=")
	fmt.Scan(&Cz)
	fmt.Print("Czmax=")
	fmt.Scan(&Czmax)
	fmt.Print("Cx=")
	fmt.Scan(&Cx)
	return
}

func DispRes(x []float64, rho, S, P, G, Cx0, k, eta, Cz, Cx, v, gamma, Vmax, Gammamax float64) { //pentru afisarea celor 4 parametri de interes
	fmt.Println("v=", v)
	fmt.Println("gamma=", gamma)
	fmt.Println("Vmax=", Vmax)
	fmt.Println("Gammamax=", Gammamax)
}
