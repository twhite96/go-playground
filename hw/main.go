package main

import "fmt"


// Use a constant
// const (
//   message = "The answer to life is %d\n"
//   answer = 42
// )

func main() {

  // declare and define type with :=
  message := "The answer to life is %d\n"

  answer := 42

	fmt.Printf(message, answer)
}

// strings === roons or unicode characters
// need to explicitly specify the use of bytes in strings
