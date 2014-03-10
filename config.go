package main

import (
  "encoding/json"
  "os"
)

type Configuration struct {
  Host    string
}

func (config *Configuration) GetConfig() (error) {
  file, err := os.Open("conf.json")
  if err != nil {
    return err
  }
  decoder := json.NewDecoder(file)
  decoder.Decode(config)

  return nil
}
