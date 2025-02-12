package pokeapi

type NamedData struct {
  Name string `json:"name"`
  URL string `json:"url"`
}

type LocationsData struct {
  Next *string `json:"next"`
  Previous *string `json:"previous"`
  Results []LocationData `json:"results"`
}

type LocationData struct {    
  Encounters []PokemonEncounters `json:"pokemon_encounters"`
  NamedData
}

type PokemonEncounters struct {
  Pokemon PokemonData `json:"pokemon"`
}

type PokemonData struct {
  BaseExperience int `json:"base_experience"`
  Height int `json:"height"`
  Weight int `json:"weight"`
  Stats []StatData `json:"stats"`
  Types []TypeData `json:"types"`
  NamedData
}

type StatData struct{
  BaseValue int `json:"base_stat"`
  Stat NamedData `json:"stat"`
}

type TypeData struct{
  Type NamedData `json:"type"`
}

