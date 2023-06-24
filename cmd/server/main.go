package main

import "fmt"

// Run() is going to be responsible for the instantiation and start up of
// our go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST api course!!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
