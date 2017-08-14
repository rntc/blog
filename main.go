package main

import (
	"log"
	"os"
	"blog/search"
	//rss.go file contains the init function to register RSS to be used. Importing
	//	will call the init method in the rss.go file.
	_"blog/matchers"

)

//init is called prior the main
func init(){
	//change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

//main is the entry point of the program
func main(){
	//perform the search for the specified term.
	search.Run("president")
}