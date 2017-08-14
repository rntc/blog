package main

import "fmt"

const name = "Rafael"

func main(){
	aName := "Reis"
	fmt.Println(greetings(name))
	fmt.Println(aName)
}

func greetings(str string) string{
	return "Welcome " + str
}
