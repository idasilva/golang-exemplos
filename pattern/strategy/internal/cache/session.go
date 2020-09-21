package cache

import "log"

type Session struct {
}

func (r *Session) push(key, value string, ttl int64) error {
	log.Println("File")
	return nil

}

func (r *Session) pop(key string) error {
	log.Println("File")
	return nil
}
