package main

import "strconv"

// Set implementation using a hash map
type Set struct {
	list map[interface{}]struct{}
}

// NewEmptySet creates an empty Set
func NewEmptySet() *Set {
	newSet := &Set{}
	newSet.list = make(map[interface{}]struct{})
	return newSet
}

// NewSet creates and initializes a Set
func NewSet(data ...interface{}) *Set {
	s := NewEmptySet()

	for _, arg := range data {
		switch v := arg.(type) {
		case []int:
			for _, element := range v {
				s.Add(element)
			}
		case []string:
			for _, element := range v {
				s.Add(element)
			}
		case int:
			s.Add(v)
		case string:
			s.Add(v)
		default:
			panic("Unsupported datatype")
		}
	}

	return s
}

// Add inserts new value/s into set
func (set *Set) Add(arr ...interface{}) {
	for _, v := range arr {
		set.list[v] = struct{}{}
	}
}

// Intersects checks if two sets intersect
func (set *Set) Intersects(other *Set) bool {
	for v := range set.list {
		_, ok := other.list[v]
		if !ok {
			return false
		}
	}
	return true
}

// Intersection returns the intersection set
func (set *Set) Intersection(other *Set) *Set {
	res := NewSet()

	for v := range set.list {
		_, ok := other.list[v]
		if !ok {
			continue
		}
		res.list[v] = struct{}{}
	}

	return res
}

func (set *Set) String() string {
	str := []rune("Set{")

	for k := range set.list {
		switch v := k.(type) {
		case int:
			str = append(str, []rune(strconv.Itoa(v))...)
		case string:
			str = append(str, []rune(strconv.Quote(v))...)
		}
		str = append(str, []rune(", ")...)
	}

	str[len(str)-2] = rune('}')

	return string(str)
}
