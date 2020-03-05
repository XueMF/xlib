package xlib

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Uniform is a type that generates uniform random numbers with the given bounds
type Uniform struct {
	Min float64
	Max float64
}

// Normal is a type that generates normal random numbers with the given bounds
type Normal struct {
	Min float64
	Max float64
}

// Distribution interface define random number generation
type Distribution interface {
	Rand() float64
}

func getRandomUint32() uint32 {
	x := time.Now().UnixNano()
	return uint32((x >> 32) ^ x)
}

// Rand is a method on Uniform
func (u Uniform) Rand() float64 {
	return rand.Float64()*(u.Max-u.Min) + u.Min
	// return u.Max - u.Min
}

// Rand is a method on Normal
func (n Normal) Rand() float64 {
	return rand.Float64()*(n.Max-n.Min) + n.Min
}

// MonteCarlo estimates the expected value of f under the distribution d using
// given number of samples. Note that f is of function type.
func MonteCarlo(f func([]float64) float64, d Distribution, samples, dim int, ch chan<- float64) {
	start := time.Now()
	var sum float64
	x := make([]float64, dim)
	for i := 0; i < samples; i++ {
		for j := range x {
			x[j] = d.Rand()
		}
		sum += f(x)
	}
	t := time.Since(start).Seconds()
	fmt.Printf("Monte: %.2fs elapsed\n", t)
	ch <- sum / float64(samples)
}

// ParallelMonteCarlo are ..
func ParallelMonteCarlo(f func([]float64) float64, d Distribution, samples, dim int) float64 {
	var ev float64
	ch := make(chan float64, 1)
	threadNumber := runtime.GOMAXPROCS(0) / 16
	sz := samples / threadNumber
	fmt.Println("sz:", sz)
	fmt.Println("threadNumber:", threadNumber)
	for i := 0; i < threadNumber; i++ {
		go MonteCarlo(f, d, sz, dim, ch)
	}
	for i := 0; i < threadNumber; i++ {
		ev += <-ch
	}
	return ev / float64(threadNumber)
}

// func doJob1(para string) {
// 	// p := Uniform{Max: -2.0, Min: 2.0}
// 	// pn := Normal{Max: -2.0, Min: 2.0}
// 	// samples, dim := xtool.Atoi(para), 3
// 	// fmt.Println("samples:", samples)
// 	// Pass the type methods as the function.
// 	// fmt.Println("Rosenbrock EV:", MonteCarlo(functions.ExtendedRosenbrock{}.Func, p, samples, dim))
// 	// fmt.Println("Beale EV:", MonteCarlo(functions.Linear{}.Func, pn, samples, dim))
// }
// func doJob2(para string) {
// 	rosenbrock := functions.ExtendedRosenbrock{}
// 	p := Uniform{
// 		Max: -2.0,
// 		Min: 2.0,
// 	}
// 	n := 100
// 	dim := 3
// 	var sum float64
// 	x := make([]float64, dim)
// 	for i := 0; i < n; i++ {
// 		for j := range x {
// 			x[j] = p.Rand()
// 			// fmt.Println(x[j])
// 		}
// 		fmt.Println(rosenbrock.Func(x))
// 		sum += rosenbrock.Func(x)
// 	}
// 	fmt.Println("Expected value is ", sum/float64(n))
// }

// func doJob3(para string) {
// 	f := functions.ExtendedRosenbrock{}.Func
// 	p := Uniform{Max: -2.0, Min: 2.0}
// 	paraInt, _ := strconv.Atoi(para)
// 	ev := ParallelMonteCarlo(f, p, paraInt, 3)
// 	fmt.Println("Expected value is:", ev)
// }
// func usage(info string) {
// 	if len(info) > 1 {
// 		fmt.Fprintf(os.Stderr, "%s\n", info)
// 	}
// 	fmt.Fprintf(os.Stderr, "--------------  %s  --------------\n", _programName)
// 	fmt.Fprintf(os.Stderr, "%s\n", _programDescription)
// 	fmt.Fprintf(os.Stderr, "%s %s <%s>\n", _programDate, _programVersion, _programAuthor)
// 	fmt.Fprintf(os.Stderr, "--------------  %s  --------------\n", _programName)
// 	fmt.Fprintf(os.Stderr, "Usage: %s 1000\n", _programName)
// 	os.Exit(1)
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	if len(os.Args) <= 1 {
// 		usage("need integer!")
// 	} else if len(os.Args) == 2 {
// 		// doJob1(os.Args[1])
// 		// doJob2(os.Args[1])
// 		start := time.Now()
// 		doJob3(os.Args[1])
// 		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// 	} else {
// 		usage("wrong command!")
// 	}
// }
