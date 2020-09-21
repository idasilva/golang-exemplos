package cache

import "log"

type Cache struct {
	Strategy strategy
}

func (c *Cache) Push(key,value string,ttl int64) error {
	log.Println("Cache")
	return c.Strategy.push(key,value,ttl)

}

func(c *Cache) Pop(key string)error{
	log.Println("Cache")
	return c.Strategy.pop(key)
}
