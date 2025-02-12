package main

import(
  "fmt"
  "os"
  "internal/pokeapi"
  "math/rand"
)

type cliCommand struct {
  Name string
  Description string
  Callback func(*config, []string) error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand {
    "exit": {
      Name: "exit",
      Description: "Exit the Pokedex",
      Callback: commandExit,
    },
    "help": {
      Name: "help",
      Description: "Displays a help message",
      Callback: commandHelp,
    },
    "map": {
      Name: "map",
      Description: "Get the next 20 locations",
      Callback: commandMap,
    },
    "mapb": {
      Name: "mapb",
      Description: "Get the previous 20 locations",
      Callback: commandMapBack,
    },
    "explore": {
      Name: "explore <location>",
      Description: "Get all of the Pokemon in a location",
      Callback: commandExplore,
    },
    "catch": {
      Name: "catch <name>",
      Description: "Attempt to catch a Pokemon",
      Callback: commandCatch,
    },
    "inspect": {
      Name: "inspect <name>",
      Description: "Inspect a caught Pokemon",
      Callback: commandInspect,
    },
    "pokedex": {
      Name: "pokedex",
      Description: "Show all of your caught Pokemon",
      Callback: commandPokedex,
    },
  }
}

func commandExit(cfg *config, prms []string) error {
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
}

func commandHelp(cfg *config, prms []string) error {
  fmt.Println()
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:")
  for _, v := range(getCommands()) {
    fmt.Printf("%v: %v\n", v.Name, v.Description)
  }
  fmt.Println()
  return nil
}

func commandMap(cfg *config, prms []string) error {
  url := *cfg.NextLocationURL
  return runCommandMap(url, cfg)
}

func commandMapBack(cfg *config, prms []string) error {
  if cfg.PreviousLocationURL == nil {
    fmt.Println("you're on the first page")
    return nil
  }
  url := *cfg.PreviousLocationURL
  return runCommandMap(url, cfg)
}

func runCommandMap(url string, cfg *config) error {
  locations, err := pokeapi.GetData[pokeapi.LocationsData](url, cfg.Cache)
  if err != nil {
    return err
  }
  cfg.NextLocationURL = locations.Next
  cfg.PreviousLocationURL = locations.Previous
  for _, r := range(locations.Results) {
    fmt.Println(r.Name)
  }
  return nil
}

func commandExplore(cfg *config, prms []string) error {
  if len(prms) == 0 {
    return fmt.Errorf("Missing command parameter! Please include a location")
  }
  for _, prm := range(prms) {
    fmt.Printf("[Exploring %v...]\n", prm)
    url := baseLocationURL + prm
    location, err := pokeapi.GetData[pokeapi.LocationData](url, cfg.Cache)
    if err != nil {
      return err
    }
    for _, r := range(location.Encounters) {
      fmt.Println(r.Pokemon.Name)
    }
  }
  return nil
}

func commandCatch(cfg *config, prms []string) error {
  if len(prms) == 0 {
    return fmt.Errorf("Missing command parameter! Please include a name")
  }
  for _, prm := range(prms) {
    fmt.Printf("Throwing a Pokeball at %v...\n", prm)
    url := "https://pokeapi.co/api/v2/pokemon/" + prm
    pokemon, err := pokeapi.GetData[pokeapi.PokemonData](url, cfg.Cache)
    if err != nil {
      return err
    }
    exp := pokemon.BaseExperience
    r := rand.Intn(700)
    if r < exp {
      fmt.Printf("%v escaped! (%v vs %v)\n", prm, r, exp)
      return nil
    }
    fmt.Printf("%v was caught! (%v vs %v)\n", prm, r, exp)
    fmt.Println("You may now inspect it with the inspect command")
    cfg.Pokedex[prm] = pokemon
  }
  return nil
}

func commandInspect(cfg *config, prms []string) error {
  if len(prms) == 0 {
    return fmt.Errorf("Missing command parameter! Please include a name")
  }
  for _, prm := range(prms) {
    pokemon, ok := cfg.Pokedex[prm]
    if !ok {
      fmt.Println("You have not caught that Pokemon!")
      return nil
    }
    fmt.Printf("Name: %v\n", prm)
    fmt.Printf("Height: %v\n", pokemon.Height)
    fmt.Printf("Weight: %v\n", pokemon.Weight)
    fmt.Println("Stats:")
    for _, stat := range(pokemon.Stats) {
      fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseValue)
    }
    fmt.Println("Types:")
    for _, t := range(pokemon.Types) {
      fmt.Printf("  -%v\n", t.Type.Name)
    }
  }
  return nil
}

func commandPokedex(cfg *config, prms []string) error {
  fmt.Println("Your Pokedex:")
  if len(cfg.Pokedex) == 0 {
    fmt.Println(" Is empty :( try catching a Pokemon with the catch command!")
    return nil
  }
  for k, _ := range(cfg.Pokedex) {
    fmt.Printf("  -%v\n", k)
  }
  return nil
}
