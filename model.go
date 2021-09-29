package euler

import (
	"fmt"
	"math"
)

const gravity = 9.8

type Ref struct {
	// Relative Height [m]
	Z float64
	// Relative Pressure [Pa]
	P float64
}

func New(r Ref, f Fluid) *Model {
	p := Pipe{start: r}
	m := Model{pipes: []Pipe{p}, fluid: f}
	return &m
}

type Model struct {
	pipes []Pipe
	fluid Fluid
}

func (m *Model) Connect(p ...Pipe) *Model {
	m.pipes = append(m.pipes, p...)
	return m
}

func (m *Model) End(P float64) (Q float64) {
	pa := m.pipes[0].start.P
	za := m.pipes[0].start.Z
	zb := m.pipes[len(m.pipes)-1].start.Z
	As := (P-pa)/(gravity*m.fluid.rho) + zb - za
	// Pipe areas

	Q = 0.01
	Dq := 10.
	i := 0
	for math.Abs(Dq/Q) > 1e-3 && i < 1000 {
		i++
		Bs := m.bs(Q)
		Qnew := math.Sqrt(-As / Bs)
		if math.IsNaN(Qnew) {
			panic("got NaN")
		}
		Dq = Q - Qnew
		Q = Qnew
		fmt.Printf("Dq=%.4g, Qnew=%g\n", Dq, Qnew)
	}

	return Q
}
func (m *Model) bs(Q float64) float64 {
	Aa := math.Pi * m.pipes[0].d * m.pipes[0].d / 4
	Ab := math.Pi * m.pipes[len(m.pipes)-1].d * m.pipes[len(m.pipes)-1].d / 4

	αa := 1.
	αb := 1.
	Bs := αb/(2*Ab*gravity) - αa/(2*Aa*gravity)
	for _, p := range m.pipes {
		A := math.Pi * p.d * p.d / 4
		u := Q / A
		Bs += p.l / p.d * estimateDarcy(p.d, u, p.ε, m.fluid.mu, m.fluid.rho)
	}
	return Q
}

type Fluid struct {
	// Pa . s
	mu float64
	// [kg/m^3]
	rho float64
}

// Estimate darcy friction factor using a variant of Colebrook equation.
func estimateDarcy(D, U, ε, mu, rho float64) float64 {
	invsqrt := -1.8 * math.Log(math.Pow(ε/(3.7*D), 1.11)+6.9/re(D, U, mu, rho))
	return 1 / (invsqrt * invsqrt)
}

// Reynolds number
func re(D, U, mu, rho float64) float64 {
	return U * rho / mu * D
}

var Water20C = Fluid{
	mu:  1e-3,
	rho: 1e3,
}
