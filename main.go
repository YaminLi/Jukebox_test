package main

import (
  "github.com/YaminLi/Jukebox_test/routes"
)

func main() {
  n  := routes.GetRouter()
  n.Run(":3000")
}
