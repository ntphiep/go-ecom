package main

import (
  "github.com/ntphiep/go-ecom/cmd/api"
  
  "log"
)

func main() {
  server := api.NewApiServer(":8080", nil)
  if err := server.Run(); err != nil {
    log.Fatal(err)
  }
}
