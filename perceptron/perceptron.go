package main

type perceptron struct {
	eta    float64
	nIter  int
	w      []float64
	errors []int
}

func (p *perceptron) fit(x [][]float64, y []int) {
	p.w = make([]float64, len(x[0])+1)
	p.errors = make([]int, p.nIter)

	for i := 0; i < p.nIter; i++ {
		errors := 0
		for j := 0; j < len(x); j++ {
			delta := p.eta * float64(y[j]-p.predict(x[j]))
			p.w[0] += delta
			for k := 0; k < len(x[0]); k++ {
				p.w[1+k] += delta * x[j][k]
			}
			if delta != 0.0 {
				errors++
			}
		}
		p.errors[i] = errors
	}
}

func (p *perceptron) predict(x []float64) int {
	if p.netInput(x) >= 0.0 {
		return 1
	} else {
		return -1
	}
}

func (p *perceptron) netInput(x []float64) float64 {
	return p.w[0] + dotProduct(x, p.w[1:])
}

func newPerceptron(eta float64, nIter int) *perceptron {
	return &perceptron{
		eta:    eta,
		nIter:  nIter,
		w:      nil,
		errors: nil,
	}
}

func dotProduct(x, y []float64) float64 {
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += x[i] * y[i]
	}
	return sum
}
