package go_koans
//
//import (
//	"fmt"
//)
//

func aboutBasics() {
	var b bool = true
	assert(b == true)  // what is truth?
	assert(b != false) // in it there is nothing false

	var i int = 1
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := 1 //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == i)
	assert(5*2 == 10)
	assert(5^2 == 7)

	var x int
	assert(x == 0) // zero values are valued in Go

	var f float32
	assert(f == 0) // for types of all types

	var s string
	assert(s == "") // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == 0)     // and types within composite types
	assert(c.f == 0) // which match the other types
	assert(c.s == "")  // in a typical way
}
