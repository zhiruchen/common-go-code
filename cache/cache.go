package cache

const (
	bucketCount = 512
)

type bucket struct {

}

// Cache thread safe in memory cache
type Cache struct {
	buckets [bucketCount]bucket
}
