package main

import (
	"basics/helpers"
	"fmt"
	"os"
	"sync"
)

func areaOfSquare(side float64) float64 {
	return side * side
}

func main() {
	fmt.Println("Hello World")
	fmt.Print("My name is Anthony Gonsalves")
	var a string = "AKELA"
	fmt.Printf("\tMai duniya mein %s hun\n", a)

	fmt.Println("Every time new line")
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("Every time same line")
	for i := 0; i < 10; i += 2 {
		fmt.Print(i, " ")
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	fmt.Println("Fprintln lets you chose the stdout for your print statement while Println used Terminal/Console for showing output")
	fmt.Fprintln(outputFile, "It prints the Output to Standard Output File by default it is Terminal/Console but in this example I have reffered the output.txt as stdout file.")

	fmt.Fprintf(outputFile, "It is similar to Fprintln but we can use Format Specifiers like %%s and %%d and replace there values with vars in respective order.")

	var age uint8 = 18
	fmt.Println("Age can not be negative so we can use it as unsigned variable")

	fmt.Println("We have self made drive package to check the age of person")

	if helpers.CheckAge(age) {
		fmt.Println("You can drive.")
	} else {
		fmt.Println("You can't drive.")
	}

	fmt.Println("Go do not have any other loop then for loop, So we use for loop to do while loop like task\nIt is similar to c++ for loop for (;a>5;) where no initial statement and update expression given\nBut in go extra ;; are also avoided")

	i := 5 // dynamically declaring and intilizing any variable but can't change the type after it
	for i > 0 {
		fmt.Println(i)
		i--
	}

	fmt.Println("Using functions from same file , which do not need any import can be declared with small first letter\nif importing any variable or function it should be declared with first letter in uppercase as I have done with CheckAge()")
	var side float64 = 5.5
	fmt.Printf("Area of square with side = %f => %f", side, areaOfSquare(side))
	fmt.Println("Formatting to only 2 decimals")
	fmt.Printf("Area of square with side = %.2f => %.2f\n", side, areaOfSquare(side))

	// Importing Person struct a self declared datatype from the helpers package

	// var p1 = helpers.Person{"Johnny Depp", 56, "Actor"}
	// above is keyless decalaration if you are confirm in the order below is decalaration with keys
	var p1 = helpers.Person{Name: "Shah Rukh Khan", Age: 60, Job: "Actor"}

	fmt.Printf("Name : %s , Age : %d , Job : %s", p1.Name, p1.Age, p1.Job)

	//Creating a method for struct which will work only with Person struct not function it will a method,
	fmt.Println("We created Unpack method because like Python and Js , go do not have built-in destructing operation")
	fmt.Println(p1)
	fmt.Println(p1.Unpack())
	name, age, job := p1.Unpack()
	fmt.Printf("After Unpacking Name :%s, Age = %d , Job = %s\n", name, age, job)
	fmt.Println(p1.Describe())
	p1.UpdateJob("Racer")
	fmt.Println(p1.Describe())

	c := helpers.Circle{Radius: 4}
	helpers.PrintArea(c)
	s := helpers.Square{Side: 4}
	helpers.PrintArea(s)

	// Concurrency Example
	var wg sync.WaitGroup

	wg.Add(2)
	go helpers.PrintLetters(&wg)
	go helpers.PrintNumber(&wg)

	wg.Wait()

}
	