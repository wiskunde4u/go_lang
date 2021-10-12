package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	pi     = 3.1415   // een constante buiten main die nu overal bruikbaar is.
	second = "second" // een string constante
	iota01 = iota
	iota02 = iota
	iota03 = iota + 6
	iota04 = 2 << iota // de iota waarde zou nu 5 moeten zijn en 2 wordt 5 posities gebitshift
	iota05 = iota + 10
	iota06 // GO neemt nu de laatste iota expressie en telt er 1 bij op dus 17
	// iota geeft dus mogelijkheden om constanten te evolven gedurende een constante block definitie..  maar blijven constanten
)
const (
	iota07 = iota // je zult zien dat de iota count terug op 0 wordt gezet bij een nieuw constante block
)

func main() {

	//   CLEAR SCREEN
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	//   END CLEAR SCREEN

	fmt.Println("Hello all")

	var i int //  compiler wil liever een waarde tijdens declaratie van daar de streepjes
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	FirstName := "Arthur"
	fmt.Println(FirstName)

	//	var f2 float32 = 3.14 //will give statement f declared but not used

	b := true
	fmt.Println(b)

	c1 := complex(3, 4) //complex imaginaire getallen  wat als ouput 3 + 4i geeft
	fmt.Println(c1)

	r, im := real(c1), imag(c1)
	fmt.Println("real = ", r, "imaginaire getal is = ", im)
	fmt.Println()
	// POINTER DATA TYPE

	var SecondName *string = new(string) //pointer variabele string
	*SecondName = "Arthur"               // dereffernce a pointer
	fmt.Println(SecondName)              // gives an memory address
	fmt.Println(*SecondName)             // geeft de inhoud van het memory address

	ThirdName := "pietjepuk"
	fmt.Println(ThirdName)
	ptr := &ThirdName // address of opperator
	fmt.Println("pointer notatie", ptr, *ptr)

	ThirdName = "kareltje"
	fmt.Println("pointer address en in houd vna de pineter", ptr, *ptr)
	fmt.Println()
	// constanten kunnen niet veranderen en in een functie wordt de contante pas op run time in gevuld niet op compile tijd
	// verschil tussen implicite en explicite
	const e = 3
	fmt.Println(e + 3)                 // c is an integer
	fmt.Println(e + 1.2)               //  c is nu een float
	const ffload int = 4               //  we zeggen nu expliciet dat d een int is
	fmt.Println(float32(ffload) + 1.2) //  we moeten dus het type van D overschrijven om er een float bij op te kunnen tellen.

	// maar GO kan meer met constanten.  iota   en Constant Expression  maar doorvoor moet de constant in de package level.

	fmt.Println("de waarde van de constante pi = ", pi)
	fmt.Println("de waarde van de constante second = ", second)
	fmt.Println(iota01, iota02) // iota constanten  zijnde de 2 constante gedefinieerd, tellende bij 0  en verhoogt zich steeds met 1
	// hier zijn iota de 2 de en 3de constante, te starten bij 0
	// met deze eigenschap kunnen we dus ook iota expressie maken
	fmt.Println("waarde van iota + 6 zou dus 4 + 6 moeten zijn", iota03)
	fmt.Println("waarde van iota=2  bit-shifted 5 =", iota04)
	fmt.Println("CHECK:  2 << 5 =", 2<<5)
	fmt.Println("iota06 zou dus 6 + 10 + 1 moeten zijn=", iota06)
	fmt.Println("iota07 moet dus weer 0 zijn =>", iota07)
	fmt.Println()

	//  Collections
	//  ARRAY is the simple fix sized collection of simular data types, like int or bool or float..
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println("arr1 output likes ", arr1)
	fmt.Println("arr1 second element ", arr1[1])
	fmt.Println()

	// reduced array decleration with implicite initalization syntax
	arr2 := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println("arr2 output likes ", arr2)
	fmt.Println("arr2 second element ", arr2[1])
	fmt.Println()

	// more flexible array  dynamically sized  called  SLIZED DATA TYPE
	// create slice array from begin to end of the array
	slice1 := arr2[:]
	fmt.Println("arr2 output is now ", arr2)
	fmt.Println("slice array output  ", slice1)
	fmt.Println()

	// NOW KEEP IN MIND   slice is value typing  so if we alter arr2 item and slice item,,  it both will be updated
	// slice1 is a kind of a pointer pointing to the arr2  it is not really a pointer.    You can do this when work with an arry comming from outside like API
	arr2[1] = 99
	slice1[2] = 101
	fmt.Println("arr2 output altered to ", arr2)
	fmt.Println("slice array output  ", slice1)
	fmt.Println()

	// SLICE can also be initialized without specific SIZE   compiler should manage the underlining array and we are not doing this...
	slice2 := []int{0, 1, 2}
	fmt.Println("slice array output  ", slice2)
	fmt.Println()

	//  Now lets grow the slice with another value which is not possible wiht static arrays

	slice2 = append(slice2, 9, 11, 33, 44, 55)
	// Go will recognize and create a new array move the slice to point to new array with data
	fmt.Println("appended slice array output  ", slice2)
	fmt.Println()

	slice3 := slice2[1:]  // start with  index element [1] and the rest
	slice4 := slice2[:2]  // copies index element [0,] and [1]  NOT [2]
	slice5 := slice2[1:3] //  Copies index element [1] and [2] NOT [3]
	fmt.Println("slice2 output  ", slice2)
	fmt.Println("slice3 output  ", slice3)
	fmt.Println("slice4 output  ", slice4)
	fmt.Println("slice5 output  ", slice5)
	fmt.Println()

	// WORKING WIHT MAPS collecton types associate arbitrary keys with  values that we are storing in the collection, needs some more carefull declaration
	map1 := map[string]int{"foo": 42} // this is a map with ONE elemenet and ONE Value, the key is FOO the value is 42
	fmt.Println("map1 output =", map1)
	fmt.Println("print just an element of the map[foo] =", map1["foo"]) //value 42 printed
	map1["foo"] = 999
	fmt.Println("print new value of map[foo] =", map1["foo"])

	// DELETE ELEMENTS FROM THE MAP   delete has 2 parmaters  frist the name of the map second the name of the element
	delete(map1, "foo")
	fmt.Println("map1 output =", map1)
	map1["karel"] = 20
	fmt.Println("map1 output =", map1)
	fmt.Println("aantal elementen in   map1", len(map1)) // is dus 1 alleen karel
	fmt.Println()
	/*
			ook een leuk voorbeeld
			var employee = make(map[string]int)
		    employee["Mark"] = 10
		    employee["Sandy"] = 20
	*/

	// STRUCT DATA TYPE   array and slice all values have to be of the same type..
	// For a map all  KEYS have to be of the same time and all values are have to be of the same type.
	// STRUCT we need to define before using it.
	type struct1 struct {
		ID        int
		FirstName string
		LastName  string
		age       int
	}
	var user struct1
	fmt.Println("user looks like", user) // user is initialized with zero values..  for integers are 0  strings are empty
	user.FirstName = "Arthur"
	user.LastName = "Dent"
	user.age = 33
	user.ID = 11
	fmt.Println("user looks like", user)
	fmt.Println("First Name", user.FirstName)
	fmt.Println("Last Name", user.LastName)

	user2 := struct1{ID: 2, FirstName: "Pietje", LastName: "Puk", age: 16}
	fmt.Println("user2 looks like", user2)

	// TO MAKE multi line initializer easier do it like following but wathc  the  COMMA  at the end declaration line.
	// Strucs kunnen tijden package worden gemaakt.
	user3 := struct1{
		ID:        3,
		FirstName: "karel",
		LastName:  "pukkel",
		age:       28, // LET OP DE KOMMA HIER !!!!
	}
	fmt.Println("user3 looks like", user3)
	fmt.Println()

}
