package datastructure

// HashMap basic hash table implementation
type HashMap interface {
	Get(string) interface{}
	Set(string, interface{})
	Hash(string) int
}

// ErrBucket if bucket cannot be found
type ErrBucket string

func (err ErrBucket) Error() string {
	return "Bucket not found"
}

type bucket struct {
	key   string
	value interface{}
	next  *bucket
}

// Map implements HashMap
type Map struct {
	buckets []*bucket
	len     int
}

// NewMap returns a new map
func NewMap(length int) *Map {
	m := &Map{make([]*bucket, length), length}
	return m
}

// Get returns value from key
func (m *Map) Get(key string) (interface{}, error) {
	hash := m.Hash(key)

	b := m.buckets[hash]

	if b != nil {
		for b.key != key {
			if b.next == nil {
				break
			}
			b = b.next
		}
		return b.value, nil
	}

	return nil, ErrBucket("")
}

// Set sets a new key-value pair
func (m *Map) Set(key string, value interface{}) {
	hash := m.Hash(key)

	b := m.buckets[hash]
	if b != nil {
		for b.next != nil {
			b = b.next
		}
		b.next = &bucket{key, value, nil}
	} else {
		m.buckets[hash] = &bucket{key, value, nil}
	}
}

// Hash returns hash of string
func (m *Map) Hash(str string) int {
	hash := 0

	for i := 0; i < len(str); i++ {
		hash += int(str[i])
	}

	return hash % m.len
}
