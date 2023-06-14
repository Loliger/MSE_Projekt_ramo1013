package src

func Swap[T any](x *T, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

// gerneric translation
func SwapGT(x *interface{}, y *interface{}) {
	tmp := *x
	*x = *y
	*y = tmp
}

// monomorphization
func SwapMMInt(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func SwapMMBool(x *bool, y *bool) {
	tmp := *x
	*x = *y
	*y = tmp
}

func SwapMMString(x *string, y *string) {
	tmp := *x
	*x = *y
	*y = tmp
}
