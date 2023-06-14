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

	fmt.Printf("Non-Generic Sums (gerneric translation): %v and %v\n",
		src.SumGT(ints),
		src.SumGT(floats))

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

	fmt.Printf("Generic:\n")
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", a1, a2)
	src.Swap(&a1, &a2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", a1, a2)
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", b1, b2)
	src.Swap(&b1, &b2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", b1, b2)
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", c1, c2)
	src.Swap(&c1, &c2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", c1, c2)

	fmt.Printf("Non-Generic (gerneric translation):\n")
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", a1, a2)
	src.SwapGT(&a1, &a2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", a1, a2)
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", b1, b2)
	src.SwapGT(&b1, &b2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", b1, b2)
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", c1, c2)
	src.SwapGT(&c1, &c2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", c1, c2)

	fmt.Printf("Non-Generic (monomorphization):\n")
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", a1, a2)
	src.SwapMMInt(&a1, &a2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", a1, a2)
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", b1, b2)
	src.SwapMMBool(&b1, &b2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", b1, b2)
	fmt.Printf("Swap Int: Before : a1=%v and a2=%v\n", c1, c2)
	src.SwapMMString(&c1, &c2)
	fmt.Printf("Swap Int: After : a1=%v and a2=%v\n", c1, c2)
}

func main() {
	nodeResults()
	sumResults()
	swapResults()
}
