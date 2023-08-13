package main

import (
	"fmt"

	"github.com/lostinopesrc/golang-play/conditions"
	"github.com/lostinopesrc/golang-play/greeting"
	"github.com/lostinopesrc/golang-play/sum"
)

func main() {
	fmt.Println("Hello this is the Entry Point")
	fmt.Println(greeting.Greeting())
	fmt.Println("THe result from sum package is:", sum.Sum())
	conditions.Conditions()
}
