package main

import (
	"algodaily/questions"
	"fmt"
)

func main() {
	//Day 1: Intersection between two arrays
	set1 := questions.NewSet([]int{2, 4, 6})
	set2 := questions.NewSet([]int{2, 4, 9, 20})
	fmt.Println(set1.Intersection(set2))
	fmt.Println()

	//Day 2: Are two strings anagrams
	fmt.Println("Mary and Army anagrams? ", questions.IsAnagram("Mary", "Army"))
	fmt.Println("Cinema and Iceman anagrams? ", questions.IsAnagram("Cinema", "Iceman"))
	fmt.Println("Jake and Jay anagrams? ", questions.IsAnagram("Jake", "Jay"))
	fmt.Println()

	//Day 3: Is string a palindrome
	fmt.Println("Is racecar a palindrome? ", questions.Palindrome("racecar"))
	fmt.Println("Is burger a palindrome? ", questions.Palindrome("burger"))
	fmt.Println()

	//Day 4: Reverse alphabetic characters only
	fmt.Println("sea!$hells3 alphabetic characeters reversed: ", questions.ReverseAlphaChars("sea!$hells3"))
	fmt.Println("1kas90jda3 alphabetic characeters reversed: ", questions.ReverseAlphaChars("1kas90jda3"))
	fmt.Println()

	//Day 5: FizzBuzz
	fmt.Println("Fizzbuzz to 15: ", questions.FizzBuzz(15))
	fmt.Println()

	//Day 6: Lonely Number
	fmt.Println("Lonely number in [4, 4, 6, 1, 3, 1, 3]: ", questions.LonelyNumber([]int{4, 4, 6, 1, 3, 1, 3}))
	fmt.Println("Lonely number in [3, 3, 9]: ", questions.LonelyNumber([]int{3, 3, 9}))
	fmt.Println("Lonely number in [1]: ", questions.LonelyNumber([]int{1}))

	//Day 7: Power of Three
	fmt.Println("Is 9 a power of three? ", questions.PowerOfThree(9))
	fmt.Println("Is 7 a power of three? ", questions.PowerOfThree(7))
	fmt.Println("Is 729 a power of three? ", questions.PowerOfThree(729))
	fmt.Println("Is 6 a power of three? ", questions.PowerOfThree(6))
}
