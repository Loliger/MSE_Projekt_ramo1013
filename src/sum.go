package src

func Sum[T int | float32](xs []T) T {
	var x T
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x

}

// gerneric translation
func SumGT(xs []interface{}) interface{} {
	var x interface{}
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x

}

// monomorphization
func SumMMInt(xs []int) int {
	var x int
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x
}

func SumMMFloat(xs []float32) float32 {
	var x float32
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x
}
