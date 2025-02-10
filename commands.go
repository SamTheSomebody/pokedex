package main

import(
  "fmt"
  "os"
)

type cliCommand struct {
  name string
  desciption string
  callback func() error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand {
    "exit": {
      name: "exit",
      desciption: "Exit the Pokedex",
      callback: commandExit,
    },
    "help": {
      name: "help",
      desciption: "Displays a help message",
      callback: commandHelp,
    },
  }
}

func commandExit() error {
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
}

func commandHelp() error {
  fmt.Println()
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:\n")
  for _, v := range(getCommands()) {
    fmt.Printf("%v: %v\n", v.name, v.desciption)
  }
  fmt.Println()
  return nil
}
