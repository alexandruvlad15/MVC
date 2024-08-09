package Model

import (
	"math"

	"gonum.org/v1/gonum/floats"   // lucru cu vectori și matrici
	"gonum.org/v1/gonum/optimize" // optimizare (minimalizare)
)

func Ecuatii(x []float64, rho, S, P, G, Cx0, k, eta, Cz, Cx float64) []float64 {
	v := x[0]                                                                     // viteza
	gamma := x[1]                                                                 // unghiul
	F1 := P*eta - (rho / 2 * S * math.Pow(v, 3) * Cx) - (v * G * math.Sin(gamma)) // prima functie
	F2 := G*math.Cos(gamma) - (rho / 2 * S * math.Pow(v, 2) * Cz)                 // a doua functie
	return []float64{F1, F2}
}

func FunctieObiectiv(x []float64, rho, S, P, G, Cx0, k, eta, Cz, Cx float64) float64 {
	f := Ecuatii(x, rho, S, P, G, Cx0, k, eta, Cz, Cx) // aflăm cele 2 funcții
	return floats.Norm(f, 2)                           // calculăm sqrt(F1^2+F2^2), astfel ecuațiile F1=0 și F2=0 se reduc doar la una
}

func Optimize(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax float64) []float64 {
	problem := optimize.Problem{
		Func: func(x []float64) float64 {
			return FunctieObiectiv(x, rho, S, P, G, Cx0, k, eta, Cz, Cx)
		},
	}
	start := []float64{50, 0.01} // valori inițiale pentru v și gamma
	result, _ := optimize.Minimize(problem, start, nil, nil)
	return result.X
}

func CalcVmax(G, rho, S, Czmax float64) float64 {
	return math.Sqrt((2 * G) / (rho * S * Czmax))
}

func CalcGammamax(P, eta, rho, S, Vmax, Cx0, k, Czmax, G float64) float64 {
	argAsin := (P*eta - (rho/2)*S*math.Pow(Vmax, 3)*(Cx0+k*math.Pow(Czmax, 2))) / (Vmax * G)
	return math.Asin(argAsin)
}
