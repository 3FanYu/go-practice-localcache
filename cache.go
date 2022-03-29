package localcache

type Cache interface {
	Set(k string, v interface{}) error
	Get(k string) (interface{}, error)
}
