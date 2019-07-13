package main

import (
	"fmt"
)

// Change changes stuff
func Change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	//Day 1: Intersection between two arrays
	fmt.Printf("Hello, World!\n")
	Intersection([]int{2, 4, 6}, []int{2, 4, 9, 20})
	fmt.Println("\u03A9")

	//Dunno
	welcome := []string{"hello", "world"}
	Change(welcome...)
	fmt.Println(welcome)

	//Day 2: Are two strings anagrams
	fmt.Println("Mary and Army anagrams? ", IsAnagram("Mary", "Army"))
	fmt.Println("Cinema and Iceman anagrams? ", IsAnagram("Cinema", "Iceman"))
	fmt.Println("Jake and Jay anagrams? ", IsAnagram("Jake", "Jay"))

	//Day 3: Is string a palindrome
	fmt.Println("Is racecar a palindrome? ", Palindrome("racecar"))
	fmt.Println("Is burger a palindrome? ", Palindrome("burger"))

	//Day 4
}
