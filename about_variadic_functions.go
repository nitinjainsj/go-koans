package go_koans

import "strings"

func concatNames(sep string, names ...string) string {
	return strings.Join(names, sep) // variadic parameters are really just slices
}

func aboutVariadicFunctions() {
	{
		str := concatNames(" ", "bob", "billy", "fred")
		assert(str == "bob billy fred") // several values can be passed to variadic parameters
	}

	// NJ: Observer the "..." in passing the arguments. This is needed since we
	//		 are passing an array. it needs to be converted to separate args of
	//		 string
	{
		names := []string{"bob", "billy", "fred"}
		str := concatNames("-", names...)
		assert(str == "bob-billy-fred") // or a slice can be dotted in place of all of them
	}
}
