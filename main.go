package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

func main() {
  cliCommand := getCommands()
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    input := scanner.Text()
    commands := cleanInput(input)

    for _, x := range(commands) {
      c, ok := cliCommand[x]
      if !ok {
        fmt.Print("Unknown command")
        continue
      }
      c.callback()
    }
  }
}

func cleanInput(text string) []string {
  x := strings.ToLower(text)
  output := strings.Fields(x)
  return output
}
