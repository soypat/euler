package euler

type Pipe struct {
	// Pipe starting Altitude [m]
	z float64
	// Interior diameter [m]
	d float64
	// Length [m]
	l float64
	// Interior relative surface roughness
	ε float64
}

func NewBasicPipe(D, ε float64) PipeType {
	return PipeType{
		kind: makePipe(D, 0, 0, ε),
	}
}

type PipeType struct {
	kind Pipe
}

func (p PipeType) New(L, z float64) Pipe {
	return makePipe(p.kind.d, z, L, p.kind.ε)
}

func makePipe(D, altitude, length, ε float64) Pipe {
	return Pipe{
		z: altitude,
		ε: ε,
		l: length,
		d: D,
	}
}
