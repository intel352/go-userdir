package main

import (
  "fmt"
  "github.com/wilmoore/go-userdir"
)

func main() {
  fmt.Println("Home Directory:",  userdir.Home())
  fmt.Println("Authorized Keys:", userdir.AuthorizedKeys())
}

