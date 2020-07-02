package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email string
	zip   string
}

func main() {
	wayne := person{firstName: "Wayne", lastName: "Andrew"}
	wayne.contactInfo = contactInfo{email: "wayne.andrew@gmail.com", zip: "98010"}
	erica := new(person) // returns a pointer
	erica.lastName = "Andrew"
	erica.firstName = "Erica"
	var nicholas person
	nicholas.firstName = "Nicholas"
	nicholas.lastName = "Andrew"
	fmt.Println(wayne)
	wayne.print()
	fmt.Println(erica)
	erica.print()
	erica.lastName = "Hegabarth"
	erica.print()
	fmt.Println(nicholas)
	nicholas.print()

	lincoln := person{
		firstName: "Lincoln",
		lastName:  "Andrew",
		contactInfo: contactInfo{
			email: "lincoln@email.com",
			zip:   "98010",
		},
	}
	lincoln.print()
	// & (operator) = give me the address of the value this variable is pointing at
	// lincoln pointer becomes an address
	//lincolnPointer := &lincoln
	//lincolnPointer.updateName("Linc")
	// because the reciever is a pointer type it will automatically dereference it.
	lincoln.updateName("Linc")
	lincoln.print()

	mySlice := []string{"Hi", "There", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

// * = on a type is a pointer to a type that is, it describes we are looking for a memory address to a value of a person
// type in this case
func (p *person) updateName(fName string) {
	// * (operator) = give me the value that is written to this memory address
	//(*p).firstName = fName <- don't know that we need this???

	// shortcut - go automatically dereferences it
	p.firstName = fName
}

// ex lincolnPointer -> 0001 : person{firstName:"Lincoln"} <-lincoln
//    addr    value
// turn addr into value with *addr (* as operator)
// turn value into addr with &value (& as operator)
//
// a *in front of a type (*person) is NOT an operator.  It's JUST describing that that type is going to be a pointer not a value

// gotcha because no pointer :-D
func updateSlice(s []string) {
	s[0] = "Bye"
}
