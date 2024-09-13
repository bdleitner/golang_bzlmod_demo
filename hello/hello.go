package main

import (
  "fmt"

  "bdl.com/greetings"
)

func main() {
  // Get a greeting message and print it.
  message := greetings.Hello("Gladys")
  fmt.Println(message)
}
