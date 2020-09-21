package cache

import "log"

type File struct {
}

func (f *File) push(key, value string, ttl int64) error {
	log.Println("File")
	return nil

}

func (f *File) pop(key string) error {
	log.Println("File")
	return nil
}
