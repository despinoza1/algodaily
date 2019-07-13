package main

import (
	"fmt"
)

func main() {
	//Day 1: Intersection between two arrays
	set1 := NewSet([]int{2, 4, 6})
	set2 := NewSet([]int{2, 4, 9, 20})
	fmt.Println(set1.Intersection(set2))
	fmt.Println()

	//Day 2: Are two strings anagrams
	fmt.Println("Mary and Army anagrams? ", IsAnagram("Mary", "Army"))
	fmt.Println("Cinema and Iceman anagrams? ", IsAnagram("Cinema", "Iceman"))
	fmt.Println("Jake and Jay anagrams? ", IsAnagram("Jake", "Jay"))
	fmt.Println()

	//Day 3: Is string a palindrome
	fmt.Println("Is racecar a palindrome? ", Palindrome("racecar"))
	fmt.Println("Is burger a palindrome? ", Palindrome("burger"))
	fmt.Println()

	//Day 4: Reverse alphabetic characters only
	fmt.Println("sea!$hells3 alphabetic characeters reversed: ", ReverseAlphaChars("sea!$hells3"))
	fmt.Println("1kas90jda3 alphabetic characeters reversed: ", ReverseAlphaChars("1kas90jda3"))
	fmt.Println()

	//Day 5: FizzBuzz
	fmt.Println("Fizzbuzz to 15: ", FizzBuzz(15))
	fmt.Println()
}
