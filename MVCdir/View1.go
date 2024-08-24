package View1

import (
	"fmt" //pentru citire si afisare
)

func Read() (float64, float64, float64, float64) { //pentru citirea datelor(parametrilor)
	var rho, v, Cz, Cx float64
	fmt.Print("rho=") //rho-densitatea aerului
	fmt.Scan(&rho)
	fmt.Print("v=") //viteza
	fmt.Scan(&v)
	fmt.Print("Cz=") //coeficientul de portanta
	fmt.Scan(&Cz)
	fmt.Print("Cx=") //coeficientul de rezistenta la inaintare
	fmt.Scan(&Cx)
	return rho, v, Cz, Cx
}

func DispRes(rho, v, Cz, Cx, P, D float64) { //pentru afisarea celor 4 parametri de interes
	fmt.Println("P=", P)
	fmt.Println("D=", D)
}
