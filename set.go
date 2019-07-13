package main

// Set implementation using a hash map
type Set struct {
	list map[int]struct{}
}

// NewSet creates and initializes a new Set
func NewSet() *Set {
	newSet := &Set{}
	newSet.list = make(map[int]struct{})
	return newSet
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
func Intersection(nums1 []int, nums2 []int) *Set {
	s1 := &Set{}
	s2 := &Set{}

	s1.list = make(map[int]struct{})
	s2.list = make(map[int]struct{})

	for _, v := range nums1 {
		s1.list[v] = struct{}{}
	}

	for _, v := range nums2 {
		s2.list[v] = struct{}{}
	}

	res := &Set{}
	res.list = make(map[int]struct{})

	for v := range s1.list {
		_, ok := s2.list[v]
		if !ok {
			continue
		}
		res.list[v] = struct{}{}
	}

	/*for v := range res.list {
		fmt.Println(v)
	} */
	return res
}
