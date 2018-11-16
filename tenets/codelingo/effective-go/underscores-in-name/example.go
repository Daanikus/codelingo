<<<<<<< HEAD
// Package main used for testing of tenet
package main

func main() {
	var a = 15
	a_bad_name := 20
	var another_bad_name = "hello"
=======
package main

func main() {
	var a int = 15
	a_bad_name := 20
	var another_bad_name string = "hello"
>>>>>>> d220f90d19daef602afee3448fc4a0f8e0f61608
	aGoodVar := 2
}

func aFunction() {
	fmt.Println("This name is fine")
}

func a_bad_function() {
	fmt.Println("Needs a name change")
}

<<<<<<< HEAD
// Bad_Interface used as example
type Bad_Interface interface {
	SomeMethod() int
}
// BetterInterface used as example
=======
type Bad_Interface interface {
	SomeMethod() int
}

>>>>>>> d220f90d19daef602afee3448fc4a0f8e0f61608
type BetterInterface interface {
	SomeMethod() int
}
