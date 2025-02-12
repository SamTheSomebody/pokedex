package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "internal/pokeapi"
  "time"
)

const baseLocationURL = "https://pokeapi.co/api/v2/location-area/"

type config struct {
  PokeApiClient pokeapi.Client
  NextLocationURL *string
  PreviousLocationURL *string
  Cache *cache
}

func main() {
  fmt.Println("[Launching Pokedex...]")

  cliCommand := getCommands()
  scanner := bufio.NewScanner(os.Stdin)

  pokeClient := pokeapi.NewClient(5 * time.Second)
  c := NewCache(10 * time.Second)
  nextLocationURL := baseLocationURL
  cfg := &config{
    PokeApiClient: pokeClient,
    NextLocationURL: &nextLocationURL,
    Cache: &c,
  }

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    input := scanner.Text()
    commands := cleanInput(input)

    if len(commands) == 0 {
      continue
    }

    c, ok := cliCommand[commands[0]]
    if !ok {
      fmt.Println("Unknown command")
      continue
    }
    
    err := c.Callback(cfg, commands[1:])
    if err != nil {
      fmt.Println(err)
    }
  }
}

func cleanInput(text string) []string {
  x := strings.ToLower(text)
  output := strings.Fields(x)
  return output
}
