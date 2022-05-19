package main

import (
	"fmt"
)

func numbers() {
  fmt.Print("\033[H\033[2J")
  var number1 int
  fmt.Println("Enter 1st Number: ")
  fmt.Scanln(&number1)

  var number2 int
  fmt.Println("Enter 2nd Number: ")
  fmt.Scanln(&number2)

  var operation int
  fmt.Println("[1] Addition")
  fmt.Println("[2] Subtraction")
  fmt.Println("[3] Multiplication")
  fmt.Println("[4] Division")
  fmt.Println("Enter a number 1-4: ")
  fmt.Scanln(&operation)
}

func main() {
  var number1 int
  fmt.Println("Enter 1st Number: ")
  fmt.Scanln(&number1)

  var number2 int
  fmt.Println("Enter 2nd Number: ")
  fmt.Scanln(&number2)

  var operation int
  fmt.Println("[1] Addition")
  fmt.Println("[2] Subtraction")
  fmt.Println("[3] Multiplication")
  fmt.Println("[4] Division")
  fmt.Println("Enter a number 1-4: ")
  fmt.Scanln(&operation)

  if operation == 1 {
    fmt.Printf("%v + %v = %v\n", number1, number2, number1 + number2)
    
    var mainMenu int
    fmt.Println("Type 0 to do another equation")
    fmt.Scanln(&mainMenu)
    if mainMenu == 0 {
      numbers()
    }
  }
  if operation == 2 {
    fmt.Printf("%v - %v = %v\n", number1, number2, number1 - number2)
    
    var mainMenu int
    fmt.Println("Type 0 to do another equation")
    fmt.Scanln(&mainMenu)
    if mainMenu == 0 {
      numbers()
    }    
  }
  if operation == 3 {
    fmt.Printf("%v x %v = %v\n", number1, number2, number1 * number2)
    
    var mainMenu int
    fmt.Println("Type 0 to do another equation")
    fmt.Scanln(&mainMenu)
    if mainMenu == 0 {
      numbers()
    }  
  }
  if operation == 4 {
    fmt.Printf("%v ÷ %v = %v\n", number1, number2, number1 / number2)
    
    var mainMenu int
    fmt.Println("Type 0 to do another equation")
    if mainMenu == 0 {
      numbers()
    }  
  }
}
