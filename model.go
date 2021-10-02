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
	p := Pipe{z: r.Z}
	m := Model{pipes: []Pipe{p}, fluid: f, startP: r.P}
	return &m
}

type Model struct {
	startP float64
	pipes  []Pipe
	fluid  Fluid
}

func (m *Model) Connect(p ...Pipe) *Model {
	m.pipes = append(m.pipes, p...)
	return m
}

func (m *Model) End(P float64) (Q float64) {
	pa := m.startP
	za := m.first().z
	zb := m.last().z
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
	Aa := math.Pi * m.first().d * m.first().d / 4
	Ab := math.Pi * m.last().d * m.last().d / 4

	// Calculate flow energy correction factor.
	Rea := re(m.first().d, Q/Aa, m.fluid.mu, m.fluid.rho)
	Reb := re(m.last().d, Q/Ab, m.fluid.mu, m.fluid.rho)
	αa := 1.
	αb := 1.
	if isLaminar(Rea) {
		αa = 2.
	}
	if isLaminar(Reb) {
		αb = 2.
	}

	Bs := αb/(2*Ab*gravity) - αa/(2*Aa*gravity)
	for _, p := range m.pipes {
		A := math.Pi * p.d * p.d / 4
		u := Q / A
		Bs += p.l / p.d * estimateDarcy(p.d, u, p.ε, m.fluid.mu, m.fluid.rho)
	}

	return Q
}

func (m *Model) first() Pipe {
	return m.pipes[0]
}

func (m *Model) last() Pipe {
	return m.pipes[len(m.pipes)-1]
}

type Fluid struct {
	// [Pa*s]
	mu float64
	// [kg/m^3]
	rho float64
}

// Estimate darcy friction factor using a variant of Colebrook equation.
func estimateDarcy(D, U, ε, mu, rho float64) float64 {
	Re := re(D, U, mu, rho)
	if isLaminar(Re) {
		// Laminar flow condition
		return 64 / Re
	}
	invsqrt := -1.8 * math.Log(math.Pow(ε/(3.7*D), 1.11)+6.9/Re)
	return 1 / (invsqrt * invsqrt)
}

func isLaminar(Re float64) bool {
	return Re < 2600
}

// Reynolds number
func re(D, U, mu, rho float64) float64 {
	return U * rho / mu * D
}

var Water20C = Fluid{
	mu:  1e-3,
	rho: 1e3,
}
