package pokeapi 

import (
  "encoding/json"
  "net/http"
  "errors"
  "fmt"
  "io"
)

func GetData[T any](url string, cache *Cache) (T, error) {
  var t T
  data, ok := cache.Get(url)
  if ok { 
    fmt.Println("[Retrieved data from cache...]")
    err := json.Unmarshal(data, &t)
    return t, err
  }
  fmt.Println("[Sending get request...]")
  fmt.Printf("[URL: %v]\n", url)
  res, err := http.Get(url)
  if err != nil {
    return t, err
  }
  defer res.Body.Close()
  fmt.Printf("[Status: %v]\n", res.StatusCode)
  if res.StatusCode != http.StatusOK {
    return t, errors.New(res.Status)
  }
  data, err = io.ReadAll(res.Body)
  if err != nil {
    return t, err
  }
  cache.Add(url, data)
  err = json.Unmarshal(data, &t)
  return t, err
}
