package playground

func Nothing()                 {}
func OnlyInputWithNames(a int) {}
func OnlyInputWithNoNames(int) {}
func OnlyOutput() int          { return 0 }

func Both(a int) int { return a }

func MethodMinus(a, b int) int {
	return a - b
}

func MethodMod(a int, b int) (int, int) {
	return a / b, a % b
}
func MethodModWithReturnVariables(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

//
////
