package go_koans

// NJ: This is new. Anonymouns functions. Assigne them to a function variable
// 		 then use the variable as a function
func aboutAnonymousFunctions() {
	{
		i := 1
		increment := func() {
			i++
		}
		increment()

		assert(i == 2) // closures function in an obvious way
	}

	{
		i := 1
		increment := func(x int) {
			x++
		}
		increment(i)

		assert(i == 1) // although anonymous functions need not always be closures
	}

	{
		double := func(x int) int { return x * 2 }

		assert(double(3) == 6) // they can do anything our hearts desire
	}
}
