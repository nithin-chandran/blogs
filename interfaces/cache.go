package interfaces

// Cache frepresents the cache functions.
type Cache interface {
	Add(key string, value string)
	Get() []string
}
