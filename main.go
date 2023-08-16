package main

import (
	"fmt"

	"github.com/lostinopesrc/golang-play/blankidentifier"
	"github.com/lostinopesrc/golang-play/conditions"
	"github.com/lostinopesrc/golang-play/functions"
	"github.com/lostinopesrc/golang-play/greeting"
	"github.com/lostinopesrc/golang-play/loops"
	"github.com/lostinopesrc/golang-play/methods"
	"github.com/lostinopesrc/golang-play/sum"
	"github.com/lostinopesrc/golang-play/switchcase"
	"github.com/lostinopesrc/golang-play/variadicfunction"
)

func main() {
	fmt.Println("Hello this is the Entry Point")
	fmt.Println(greeting.Greeting())
	fmt.Println("THe result from sum package is:", sum.Sum())
	conditions.Conditions()
	loops.Loops()
	switchcase.SwitchCase()
	fmt.Println(functions.Swap(10, 20)) // call by value
	// call by reference
	var p int = 10
	var q int = 20
	fmt.Println(functions.Swap1(&p, &q))
	// variadic function call
	fmt.Println(variadicfunction.Join_String("Hi", "Hello"))
	// using division output as blank identifier
	mul, _ := blankidentifier.Mul_Div(1, 2)
	fmt.Println("Multiplication output is:", mul)
	// method with struct type as receiver
	res := methods.Book{
		Name:   "Learn Go",
		Author: "lostinopensrc",
	}
	res.Artifact()
}
