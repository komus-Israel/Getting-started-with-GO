package main	

//	go raises an exception if package main isnt found. It doen't have to be main but the package name must be the name
//	the functin that will serve as an entry point for your applicatioin 



import "fmt"


func main() {
	fmt.Println("what is your name")
	
	name := "timi"
	age := 3.8

	fmt.Println(name, age)
}
