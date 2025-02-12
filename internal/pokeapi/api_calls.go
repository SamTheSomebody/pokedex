package pokeapi 

import (
  "net/http"
  "errors"
  "fmt"
  "io"
)

func GetRequest(url string) (*[]byte, error) {
  res, err := http.Get(url)
  if err != nil {
    s := fmt.Sprintf("Couldn't complete GET request! Error: %v", err)
    return nil, errors.New(s)
  }
  defer res.Body.Close()
  
  fmt.Printf("[Status: %v]\n", res.StatusCode)
  if res.StatusCode != http.StatusOK {
    s := fmt.Sprintf("HTTP error! Error: %v", res.Status)
    return nil, errors.New(s)
  }

  data, err := io.ReadAll(res.Body)
  return &data, err
}
