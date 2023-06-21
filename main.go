package main

import (
	"fmt"

	"github.com/Loliger/MSE_Projekt_ramo1013/src"
)

func nodeResults() {
	fmt.Printf("Generic:")
	src.TestNode()
	fmt.Printf("\nNon-Generic (gerneric translation):")
	src.TestNode_G()
	fmt.Printf("\nNon-Generic (monomorphization):")
	src.TestNode_MM()
	fmt.Printf("\n")
}

func sumResults() {
	ints := []int{34, 12}
	floats := []float32{35.98, 26.99}

	fmt.Printf("Generic Sums: %v and %v\n",
		src.Sum(ints),
		src.Sum(floats))

	valsInts := make([]interface{}, len(ints))
	for i, v := range ints {
		valsInts[i] = v
	}
	valsFloats := make([]interface{}, len(floats))
	for i, v := range floats {
		valsFloats[i] = v
	}

	fmt.Printf("Non-Generic Sums (gerneric translation): %v and %v\n",
		src.SumGT(valsInts),
		src.SumGT(valsFloats))

	fmt.Printf("Non-Generic Sums (monomorphization): %v and %v\n",
		src.SumMMInt(ints),
		src.SumMMFloat(floats))
}

func swapResults() {
	a1 := 15
	a2 := 12
	b1 := true
	b2 := false
	c1 := "abc"
	c2 := "def"
	var a3 interface{}
	a3 = 15
	var a4 interface{}
	a4 = 12
	var b3 interface{}
	b3 = true
	var b4 interface{}
	b4 = false
	var c3 interface{}
	c3 = "abc"
	var c4 interface{}
	c4 = "def"
	fmt.Printf("Generic:\n")
	fmt.Printf("Swap int: Before : a1=%v and a2=%v\n", a1, a2)
	src.Swap(&a1, &a2)
	fmt.Printf("Swap int: After : a1=%v and a2=%v\n", a1, a2)
	fmt.Printf("Swap bool: Before : b1=%v and b2=%v\n", b1, b2)
	src.Swap(&b1, &b2)
	fmt.Printf("Swap bool: After : b1=%v and b2=%v\n", b1, b2)
	fmt.Printf("Swap string: Before : c1=%v and c2=%v\n", c1, c2)
	src.Swap(&c1, &c2)
	fmt.Printf("Swap string: After : c1=%v and c2=%v\n", c1, c2)

	fmt.Printf("Non-Generic (gerneric translation):\n")
	fmt.Printf("Swap int: Before : a1=%v and a2=%v\n", a3, a4)
	src.SwapGT(&a3, &a4)
	fmt.Printf("Swap int: After : a1=%v and a2=%v\n", a3, a4)
	fmt.Printf("Swap bool: Before : b1=%v and b2=%v\n", b3, b4)
	src.SwapGT(&b3, &b4)
	fmt.Printf("Swap bool: After : b1=%v and b2=%v\n", b3, b4)
	fmt.Printf("Swap string: Before : c1=%v and c2=%v\n", c3, c4)
	src.SwapGT(&c3, &c4)
	fmt.Printf("Swap string: After : c1=%v and c2=%v\n", c3, c4)

	fmt.Printf("Non-Generic (monomorphization):\n")
	fmt.Printf("Swap int: Before : a1=%v and a2=%v\n", a1, a2)
	src.SwapMMInt(&a1, &a2)
	fmt.Printf("Swap int: After : a1=%v and a2=%v\n", a1, a2)
	fmt.Printf("Swap bool: Before : b1=%v and b2=%v\n", b1, b2)
	src.SwapMMBool(&b1, &b2)
	fmt.Printf("Swap bool: After : b1=%v and b2=%v\n", b1, b2)
	fmt.Printf("Swap string: Before : c1=%v and c2=%v\n", c1, c2)
	src.SwapMMString(&c1, &c2)
	fmt.Printf("Swap string: After : c1=%v and c2=%v\n", c1, c2)
}

func main() {
	nodeResults()
	sumResults()
	swapResults()
}
