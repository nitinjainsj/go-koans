package go_koans

func aboutStructs() {
	var bob struct {
		name string
		age  int
	}
	bob.name = "bob"
	bob.age = 30

	assert(bob.name == "bob") // structs are collections of named variables
	assert(bob.age == 30)     // each field has both setter and getter behavior

	type person struct {
		name string
		age  int
	}

	var john person
	john.name = "bob"
	john.age  = 30

	// NJ: Seems like there is native overload of operators. Here bob and john
	//	   are not compared as pointers, but their content. The comparison is
	//		 made based on the struct member names.
	//		 type person struct { n string; a int } will fail the compile
	assert(bob == john) // assuredly, bob is certainly not john.. yet
}
