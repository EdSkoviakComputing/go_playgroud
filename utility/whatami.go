package utility

func whatAmI(i interface{}) string {
	switch i.(type) {
	case bool:
		return ("I'm a bool")
	case int:
		return ("I'm an int")
	case []int:
		return ("I'm an int array")
	default:
		return ("I'm unknown type")

	}
}
