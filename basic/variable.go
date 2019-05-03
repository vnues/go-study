// package main

// import "fmt"

// func main() {
// 	var i int
// 	var f float32 
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
