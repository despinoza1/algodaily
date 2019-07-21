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
	fmt.Println()

	//Day 7: Power of Three
	fmt.Println("Is 9 a power of three? ", questions.PowerOfThree(9))
	fmt.Println("Is 7 a power of three? ", questions.PowerOfThree(7))
	fmt.Println("Is 729 a power of three? ", questions.PowerOfThree(729))
	fmt.Println("Is 6 a power of three? ", questions.PowerOfThree(6))
	fmt.Println()

	fmt.Println("Is 81 a power of three? ", questions.IsPowerOf(81, 3))
	fmt.Println("Is 513 a power of eight? ", questions.IsPowerOf(513, 8))
	fmt.Println("Is 625 a power of five? ", questions.IsPowerOf(625, 5))
	fmt.Println("Is 256 a power of two? ", questions.IsPowerOf(256, 2))
	fmt.Println()

	//Day 8: Hash Map
	names := [4]string{"Bob", "James", "John", "Maria"}

	dict := questions.NewMap(2)

	dict.Set("Bob", 3000)
	dict.Set("James", 2100.01)
	dict.Set("John", "Doe")
	dict.Set("Maria", [4]interface{}{1, 2, "Hello", 3.14})

	for _, name := range names {
		if value, err := dict.Get(name); err == nil {
			fmt.Printf("dict[%q]=%v\n", name, value)
		}
	}
	fmt.Println()

	//Day 9: Binary Tree In Order Traversal
	var tree questions.BinaryTree = questions.NewTree(8)
	tree.Add(5)
	tree.Add(17)
	tree.Add(22)
	tree.Add(6)
	tree.Add(3)
	tree.Add(1)
	tree.Add(13)
	tree.Add(7)

	fmt.Print("Tree In Order: ")
	tree.InOrder()
	fmt.Print("\nTree Post Order: ")
	tree.PostOrder()
	fmt.Print("\nTree Pre Order: ")
	tree.PreOrder()
	fmt.Println()

	//Day 10: First Nonrepeating value
	fmt.Println("First Nonrepeat of \"\": ", questions.FirstNonRepeat(""))
	fmt.Println("First Nonrepeat of \"a\": ", questions.FirstNonRepeat("a"))
	fmt.Println("First Nonrepeat of \"hello world\": ", questions.FirstNonRepeat("hello world"))
	fmt.Println("First Nonrepeat of \"asdfsdafdasfjdfsafnnunlfdffvxcvsfansd\": ", questions.FirstNonRepeat("asdfsdafdasfjdfsafnnunlfdffvxcvsfansd"))
	fmt.Println()

	//Day 11: Sum Digits
	fmt.Println("SumDigits(49): ", questions.SumDigits(49))
	fmt.Println("SumDigits(1): ", questions.SumDigits(1))
	fmt.Println("SumDigits(439230): ", questions.SumDigits(439230))
	fmt.Println()

	//Day 12: $ Deletion
	fmt.Println("['f$ec', 'ec']: ", questions.DollarSignDeletion([]string{"f$ec", "ec"}))
	fmt.Println("['ab$$', 'c$d$']: ", questions.DollarSignDeletion([]string{"ab$$", "c$d$"}))
	fmt.Println("['b$$p', '$b$p']: ", questions.DollarSignDeletion([]string{"b$$p", "$b$p"}))
	fmt.Println("['e$ec', 'ec', 'as$$ec]: ", questions.DollarSignDeletion([]string{"e$ec", "ec", "as$$ec"}))
	fmt.Println()

	//Day 13: Detect Substring
	fmt.Println("Index of flew in thepigflewwow: ", questions.DetectSubStr("thepigflewwow", "flew"))
	fmt.Println("Index of two in twocanplay: ", questions.DetectSubStr("twocanplay", "two"))
	fmt.Println("Index of pork in wherearemyshorts: ", questions.DetectSubStr("wherearemyshorts", "pork"))
	fmt.Println()
}
