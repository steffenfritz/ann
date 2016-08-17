/* artificial neuronal network*/

package ann

import "math"

// learning param
var E float64 = 0.7

// bias node returns +/- 1.0
type Biasnode float64

// type node
type Node struct {
	id    string
	in    []float64
	netin float64
	out   float64
}

// type edge
type Edge struct {
	id   string
	from string
	to   string
	in   float64
	w    float64
	out  float64
}

// sum of a float64 slice
func (n Node) Fsum(in *[]float64) float64 {
	fs := 0.0
	for _, i := range *in {
		fs += i
	}
	return fs
}

// identity
func Id(in *float64) *float64 {
	return in
}

// log function
func (n Node) Log(netin float64) float64 {
	return 1.0 / (1 + math.Pow(math.Exp(1), -netin))
}

// hebb rule
func (e Edge) Hebb(w1 float64, w2 float64) float64 {
	return E * w1 * w2
}

// simpe delta rule
func (e Edge) Delta(isout float64, expout float64, outin float64) float64 {
	delta := expout - isout
	return E * delta * outin
}
