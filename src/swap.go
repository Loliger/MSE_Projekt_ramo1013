package src

func Swap[T any](x *T, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

// gerneric translation
func SwapGT(x *interface{}, y *interface{}) {
	switch (*x).(type) {
	case int:
		_, ok := (*y).(int)
		if ok {
			tmp := *x
			*x = *y
			*y = tmp
		}
	case bool:
		_, ok := (*y).(bool)
		if ok {
			tmp := *x
			*x = *y
			*y = tmp
		}
	case string:
		_, ok := (*y).(string)
		if ok {
			tmp := *x
			*x = *y
			*y = tmp
		}
	}
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
