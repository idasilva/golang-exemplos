package main

import (
	"github.dxc.com/projects/golang-exemplos/pattern/strategy/internal/cache"
	"log"
)

func main() {

	var c = &cache.Cache{}

	c.Strategy = &cache.File{}
	if err := c.Push("key-f", "values-f", 3600); err != nil {
		log.Println("push file")

	}
	c.Strategy = &cache.Redis{}
	if err := c.Pop("key-r"); err != nil {
		log.Println("pop redis")

	}
	c.Strategy = &cache.Session{}
	if err := c.Push("key-s", "values-s", 3600); err != nil {
		log.Println("push file")

	}

}
