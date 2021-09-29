# euler
Resolution of euler flow equation for 1D steady, incompressible flow (Bernoulli Equation)

THIS IS A **WORK IN PROGRESS** do **NOT** use for calculations.


```go
    const bar = 100e3
    m := euler.New(euler.Ref{P: 60 * bar}, euler.Water20C)
	piping := euler.NewBasicPipe(20e-3, 1e-3) // some 20mm diameter pipe
	ten := piping.New(10, 0) // 10meter long piping
	m.Connect(ten, ten, ten)
	q := m.End(bar)
    fmt.Println("q=%0.3g m^3/s")
    // Output: q=8.445 m^3/s
```