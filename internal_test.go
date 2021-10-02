package euler

const (
	bar = 100e3 // 1bar is 100kPa
)

// func TestSimplePipe(t *testing.T) {
// 	Dp := 10. // small delta pressure for laminar flow condition.
// 	fluid := Water20C
// 	D := 10e-3 // 1cm diameter pipe
// 	eps := 0.1
// 	Area := math.Pi * D * D / 4
// 	As := Dp / gravity / fluid.rho
// 	BsNoLoss := 0.0

// 	frictionLossCoef := re()
// 	pipe := NewBasicPipe(D, eps)
// 	m := New(Ref{P: 0, Z: 0}, fluid)
// }
