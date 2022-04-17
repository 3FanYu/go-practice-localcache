package localcache

// localcache provides a simple key-value store
type Cache interface {
	Set(k string, v interface{}) error
	Get(k string) (interface{}, error)
}
