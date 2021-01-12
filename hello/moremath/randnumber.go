package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

// when two or more consecutive named function parameters share a type
// we could omit the type from all but the last
func add(x, y int) int {
	return x + y
}

// any number of return value
func swap(x, y string) (string, string) {
	return y, x
}

// a return statement without arguments returns the named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}


func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// package starts with a captial number
	fmt.Println(math.Pi)

	fmt.Println(add(123,123))

	// if an initalizer is present, the type can be omitted
	var c, python, java = true, false, "no!"
	var i, j = 1,2
	fmt.Println(i, j, c, python, java)

	// byte // alias for uint8
	// rune // alias for int32, represents a Unicode code point

	// need explicit type transfer
	var f float64 = float64(i)
	u := uint(i)
	fmt.Println(i,f,u)

	// const cannot be delcared with :=

	// for, the init and post statement are optional-->while
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000{
		sum += sum
	}
	fmt.Println(sum)

	// wait until other function executed
	// deferred function calls are pushed onto a stack, when a 
	// function returns, its deferred calls are executed in LIFO order
	defer fmt.Println("defer end")

	// switch: only run selected case, no break need
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// pointer
	i, j = 42, 2701
	p := &i
	fmt.Println(*p)

	// struct
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1,2}
	v_t := &v
	v_t.X = 123
	fmt.Println(v)

	// array cannot change size
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int  = primes[1:4]
	fmt.Println(s)
}