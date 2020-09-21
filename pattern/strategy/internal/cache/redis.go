package cache

import "log"

type Redis struct {
}

func (r *Redis) push(key, value string, ttl int64) error {
	log.Println("File")
	return nil

}

func (r *Redis) pop(key string) error {
	log.Println("File")
	return nil
}
