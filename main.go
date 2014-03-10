package main

import (
  "fmt"
)

func main() {
  var cfg Configuration
  err := cfg.GetConfig()
  if err != nil {
    fmt.Println(err)
  }
  renderUI(&cfg)
}

