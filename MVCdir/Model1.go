package Model1

import (
	"math" //pentru utilizarea functiilor matematice
)

func Det_pol() ([]float64, []float64, []float64) {
	c1 := []float64{0, 0.1, 0.125}                    //functia de curbura 0.125(0.8x+x^2) pe intervalul (0,0.4]
	c2 := []float64{0.01112, 0.04448, 0.0556}         //functia de curbura 0.0556(0.2+0.8x+x^2) pe intervalul (0.4,1]
	t := []float64{0, 0.1260, 0.3516, 0.2843, 0.1036} // functia de grosime aprox yt(x)=0.1260x+0.3516x^2+0.2843x^3+0.1036x^4, consideram t=1
	return c1, c2, t
}

func Det_prod(x []float64, y []float64, z []float64) ([]float64, []float64) { //realizarea produselor dintre polinoame,dorim sa calculam integralele din c1*t si c2*t
	p1 := []float64{0, 0, 0, 0, 0, 0, 0} //cate 7 elemente, deoarece t este de grad 4, iar c1 si c2 de grad 3
	p2 := []float64{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i <= 2; i++ { //c1-polinom de grad 3
		for j := 0; j <= 4; j++ { //t-polinom de grad 4
			p1[i+j] += x[i] * z[j] //determinarea primului produs
		}
	}
	for i := 0; i <= 2; i++ { //c2-polinom de grad 3
		for j := 0; j <= 4; j++ { //t-polinom de grad 4
			p2[i+j] += y[i] * z[j] //determinarea celui de-al doilea produs
		}
	}
	return p1, p2
}

func Det_int(x float64, a []float64, b []float64, c []float64) (float64, float64) { //determinarea integralelor
	var prod1 []float64
	var prod2 []float64
	prod1, prod2 = Det_prod(a, b, c)
	var s1 float64
	var s2 float64
	s1 = 0 //pentru integralele din polinomul c1*t
	s2 = 0 //pentru integralele din polinomul c2*t
	for i := 0; i <= 6; i++ {
		s1 += prod1[i] * (math.Pow(x, float64(i+1))) / float64(i+1) //determinarea valorii primitivei polinomului c1*t in punctul x
	}
	for i := 0; i <= 6; i++ {
		s2 += prod2[i] * (math.Pow(x, float64(i+1))) / float64(i+1) //determinarea valorii primitivei polinomului c2*t in punctul x
	}
	return s1, s2
}
func Det_suprafata(x []float64, y []float64, z []float64) float64 {
	var a1 float64
	var b1 float64
	var b2 float64
	var c2 float64
	a1, _ = Det_int(0, x, y, z)    //valoarea primitivei in 0
	b1, b2 = Det_int(0.4, x, y, z) //valoarea primitivei in 0.4
	_, c2 = Det_int(1, x, y, z)    //valoarea primitivei in 1
	supr := (-a1 + b1) + (c2 - b2) //determinarea suprafetei ca suma din integralele pe cele 2 parti ale profilului
	return supr
}

func Det_P(rho float64, v float64, S float64, Cz float64) float64 { //determinarea fortei de portanta
	P := 0.5 * rho * math.Pow(v, 2) * S * Cz
	return P
}
func Det_D(rho float64, v float64, S float64, Cx float64) float64 { //determinarea fortei de rezistenta la inaintare
	D := 0.5 * rho * math.Pow(v, 2) * S * Cx
	return D
}
