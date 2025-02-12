package main

import(
  "fmt"
  "os"
  "internal/pokeapi"
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

func commandMapBack(cfg *config, prms []string) error {
  if cfg.PreviousLocationURL == nil {
    fmt.Println("you're on the first page")
    return nil
  }
  url := *cfg.PreviousLocationURL
  return runCommandMap(url, cfg)
}

func commandMap(cfg *config, prms []string) error {
  url := *cfg.NextLocationURL
  return runCommandMap(url, cfg)
}

func runCommandMap(url string, cfg *config) error {
  data, ok := cfg.Cache.Get(url)
  if !ok {
    fmt.Println("[Looking up map area...]")
    fmt.Printf("[URL: %v]\n", url)
    pData, err := pokeapi.GetRequest(url) 
    if err != nil {
      return err
    }
    data = *pData
    cfg.Cache.Add(url, data)
  } else {
    fmt.Println("[Retriving map area from cache...]")
  }

  fmt.Println("[Converting JSON to structed data...]")
  locations, err := pokeapi.BytesToData[pokeapi.LocationsData](data)
  if err != nil {
    fmt.Errorf("Data:\n\n%v\n\n", data)
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
    data, ok := cfg.Cache.Get(url)
    if !ok {
      fmt.Println("[Looking up encounters...]")
      fmt.Printf("[URL: %v]\n", url)

      pData, err := pokeapi.GetRequest(url) 
      if err != nil {
        return err
      }

      data = *pData
      cfg.Cache.Add(url, data)
    }
    
    fmt.Println("[Converting JSON to structed data...]")
    location, err := pokeapi.BytesToData[pokeapi.LocationData](data)
    if err != nil {
      fmt.Errorf("Data:\n\n%v\n\n", data)
      return err
    }

    for _, r := range(location.Encounters) {
      fmt.Println(r.Pokemon.Name)
    }
  }
  return nil
}
