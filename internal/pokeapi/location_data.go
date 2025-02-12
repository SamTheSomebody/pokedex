package pokeapi

import (
  "encoding/json"
)

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
  NamedData
}

func BytesToLocationData(data []byte) (LocationsData, error) {
  locations := LocationsData{}
  err := json.Unmarshal(data, &locations)
  return locations, err
}

func BytesToData[T any](data []byte) (T, error) {
  var t T
  err := json.Unmarshal(data, &t)
  return t, err
}

func LocationDataToBytes(data LocationsData) ([]byte, error) {
  b, err := json.Marshal(data)
  return b, err
}


