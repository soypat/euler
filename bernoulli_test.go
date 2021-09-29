package euler_test

import (
	"testing"

	"github.com/soypat/euler"
)

const (
	bar = 100e3
)

func TestBernoulli(t *testing.T) {
	m := euler.New(euler.Ref{P: 60 * bar}, euler.Water20C)
	piping := euler.NewBasicPipe(20e-3, 1e-3)
	ten := piping.New(10, 0)
	m.Connect(ten, ten, ten)
	q := m.End(bar)
	t.Error(q)
}
