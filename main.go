package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
  x := strings.ToLower(text)
  output := strings.Fields(x)
  return output
}
