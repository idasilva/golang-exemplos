package cache

type strategy interface {
	push(key, value string, ttl int64) error
	pop(key string) error
}
