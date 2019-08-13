package main

import (
	"algodaily/questions"
	"algodaily/questions/datastructure"
	"fmt"
)

func main() {
	//Day 1: Intersection between two arrays
	set1 := datastructure.NewSet([]int{2, 4, 6})
	set2 := datastructure.NewSet([]int{2, 4, 9, 20})
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

	dict := datastructure.NewMap(2)

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
	var tree datastructure.BinaryTree = datastructure.NewTree(8)
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
	fmt.Println("First Nonrepeat of \"\": ",
		questions.FirstNonRepeat(""))
	fmt.Println("First Nonrepeat of \"a\": ",
		questions.FirstNonRepeat("a"))
	fmt.Println("First Nonrepeat of \"hello world\": ",
		questions.FirstNonRepeat("hello world"))
	fmt.Println("First Nonrepeat of \"asdfsdafdasfjdfsafnnunlfdffvxcvsfansd\": ",
		questions.FirstNonRepeat("asdfsdafdasfjdfsafnnunlfdffvxcvsfansd"))
	fmt.Println()

	//Day 11: Sum Digits
	fmt.Println("SumDigits(49): ", questions.SumDigits(49))
	fmt.Println("SumDigits(1): ", questions.SumDigits(1))
	fmt.Println("SumDigits(439230): ", questions.SumDigits(439230))
	fmt.Println()

	//Day 12: $ Deletion
	fmt.Println("['f$ec', 'ec']: ",
		questions.DollarSignDeletion([]string{"f$ec", "ec"}))
	fmt.Println("['ab$$', 'c$d$']: ",
		questions.DollarSignDeletion([]string{"ab$$", "c$d$"}))
	fmt.Println("['b$$p', '$b$p']: ",
		questions.DollarSignDeletion([]string{"b$$p", "$b$p"}))
	fmt.Println("['e$ec', 'ec', 'as$$ec]: ",
		questions.DollarSignDeletion([]string{"e$ec", "ec", "as$$ec"}))
	fmt.Println()

	//Day 13: Detect Substring
	fmt.Println("Index of flew in thepigflewwow: ",
		questions.DetectSubStr("thepigflewwow", "flew"))
	fmt.Println("Index of two in twocanplay: ",
		questions.DetectSubStr("twocanplay", "two"))
	fmt.Println("Index of pork in wherearemyshorts: ",
		questions.DetectSubStr("wherearemyshorts", "pork"))
	fmt.Println()

	//Day 14: Missing in Unsorted Array
	fmt.Println("Missing number from [2, 5, 1, 4, 9, 6, 3, 7]: ",
		questions.MissingInUnsorted([]int{2, 5, 1, 4, 9, 6, 3, 7}, 1, 9))
	fmt.Println()

	//Day 15: Max of Min Pairs
	fmt.Println("Max of min pairs from [3, 4, 2, 5]: ",
		questions.MaxOfMinPairs([]int{3, 4, 2, 5}))
	fmt.Println("Max of min pairs from [1, 3, 2, 6, 5, 4]: ",
		questions.MaxOfMinPairs([]int{1, 3, 2, 6, 5, 4}))
	fmt.Println()

	//Day 16: Contiguous Subarray Sum
	fmt.Println("Is 3 a contiguous subarray sum []: ",
		questions.SubArraySum([]int{}, 3))
	fmt.Println("Is 3 a contiguous subarray sum [1, 2, 3]: ",
		questions.SubArraySum([]int{1, 2, 3}, 3))
	fmt.Println("Is 47 a contiguous subarray sum [3, 6, 12, 35]: ",
		questions.SubArraySum([]int{3, 6, 12, 35}, 47))
	fmt.Println()

	//Day 17: Sum of Primes
	fmt.Println("Sum of primes less than 2: ",
		questions.SumPrimes(2))
	fmt.Println("Sum of primes less than 30: ",
		questions.SumPrimes(30))
	fmt.Println("Sum of primes less than 55: ",
		questions.SumPrimes(55))
	fmt.Println()

	//Day 18: Min in Rotated Array
	fmt.Println("Minimum number in [6, 7, 8, 0, 1, 2, 3, 4, 5]: ",
		questions.MinInRotArray([]int{6, 7, 8, 0, 1, 2, 3, 4, 5}))
	fmt.Println("Minimum number in [6, 7, 8, 9, 10, 3, 4, 5]: ",
		questions.MinInRotArray([]int{6, 7, 8, 9, 10, 3, 4, 5}))
	fmt.Println()

	//Day 19 & 20: Uniqueness of Arrays
	fmt.Println("Unique elements in [8,8,15,6,19,7,12,6,6,3,13,9,15,14,1,13,4,11,16]:\n\t",
		questions.UniqueInArray([]int{8, 8, 15, 6, 19, 7, 12, 6, 6, 3, 13, 9, 15, 14, 1, 13, 4, 11, 16}))
	fmt.Println("Unique elements in [12,7,2,20,20,2,15,20,2,10,12,1]:\n\t",
		questions.UniqueInArray([]int{12, 7, 2, 20, 20, 2, 15, 20, 2, 10, 12, 1}))
	fmt.Println("Unique elements in [6,12,5,1,4,18,10,17,10,0,1,7,6,18,11,2,15,19]:\n\t",
		questions.UniqueInArray([]int{6, 12, 5, 1, 4, 18, 10, 17, 10, 0, 1, 7, 6, 18, 11, 2, 15, 19}))
	fmt.Println("Unique elements in [9,0,11,16,19,14,7,18,10,6,0,17,12,9,12,18,0,14,17]:\n\t",
		questions.UniqueInArray([]int{9, 0, 11, 16, 19, 14, 7, 18, 10, 6, 0, 17, 12, 9, 12, 18, 0, 14, 17}))
	fmt.Println("Unique elements in [5,10,3,17,9,12,19,4,16,19,7,9,7,8,10]:\n\t",
		questions.UniqueInArray([]int{5, 10, 3, 17, 9, 12, 19, 4, 16, 19, 7, 9, 7, 8, 10}))
	fmt.Println()

	//Day 22: Treats Distribution
	fmt.Println("Treat Distribution of [2, 2, 3, 3, 4, 4]:\n\t",
		questions.TreatDistribution([]int{2, 2, 3, 3, 4, 4}))
	fmt.Println("Treat Distribution of [1, 1, 2, 4]:\n\t",
		questions.TreatDistribution([]int{1, 1, 2, 4}))
	fmt.Println("Treat Distribution of [7, 7]:\n\t",
		questions.TreatDistribution([]int{7, 7}))
	fmt.Println()

	//Day 23: Least Missing Positive Number
	fmt.Println("Least Missing Positive Number of [3, 5, -1, 1]:\n\t",
		questions.LeastPositiveNumber([]int{3, 5, -1, 1}))
	fmt.Println("Least Missing Positive Number of [5, 6, 7, 8, 9]:\n\t",
		questions.LeastPositiveNumber([]int{5, 6, 7, 8, 9}))
	fmt.Println("Least Missing Positive Number of []:\n\t",
		questions.LeastPositiveNumber([]int{}))
	fmt.Println()

	//Day 24: Product Except Self
	fmt.Println("Product Except Self of []:\n\t",
		questions.ProductExceptSelf([]int{}))
	fmt.Println("Product Except Self of [7,8,5,18,16,11,20]:\n\t",
		questions.ProductExceptSelf([]int{7, 8, 5, 18, 16, 11, 20}))
	fmt.Println("Product Except Self of [2,0,9,19,7,10,17,3,10,11,3,3,17,18,19]:\n\t",
		questions.ProductExceptSelf([]int{2, 0, 9, 19, 7, 10, 17, 3, 10, 11, 3, 3, 17, 18, 19}))
	fmt.Println("Product Except Self of [9,9,3,4,18,8,6,18,1,6,19]:\n\t",
		questions.ProductExceptSelf([]int{9, 9, 3, 4, 18, 8, 6, 18, 1, 6, 19}))
	fmt.Println()

	//Day 25: Two Sum
	fmt.Println("Indices of [1, 9, 13, 20, 47] that sum to 10:\n\t",
		questions.TwoSum([]int{1, 9, 13, 20, 47}, 10))
	fmt.Println("Indices of [3, 2, 4, 1, 9] that sum to 12:\n\t",
		questions.TwoSum([]int{3, 2, 4, 1, 9}, 12))
	fmt.Println("Indices of [] that sum to 10:\n\t",
		questions.TwoSum([]int{}, 10))
	fmt.Println()

	//Day 26: Intersection of Linked List
	list1 := datastructure.NewDoublyLinkedList(4)
	list2 := datastructure.NewDoublyLinkedList(2)

	for _, val := range []int{5, 6, 7, 8, 9, 10} {
		list1.Append(val)
	}
	for _, val := range []int{3, 4, 5, 6, 7, 8} {
		list2.Append(val)
	}

	fmt.Println("The intersection of [4, 5, 6, 7, 8, 9, 10] and [2, 3, 4, 5, 6, 7, 8] is:\n\t",
		datastructure.GetIntersection(*list1, *list2))
	fmt.Println()

	//Day 27: Get a random node from linked list
	fmt.Println("Random node from the list [4, 5, 6, 7, 8, 9, 10]:\n\t",
		datastructure.GetRandomNode(*list1))
	fmt.Println()

	//Day 28: Swap every pair of elements of linked list
	list1.SwapEveryTwo()
	fmt.Println("Every two nodes swapped from the list [4, 5, 6, 7, 8, 9, 10]:\n\t",
		list1)
	fmt.Println()

	//Day 29: Get Union of two linked list
	union := datastructure.GetUnion(*list1, *list2)
	fmt.Println("Union of two linked list:\n\t",
		union)
	fmt.Println()

	//Day 30: Delete node from linked list
	list1.Delete(6)
	fmt.Println("Linked list after deletting node with value 6:\n\t",
		list1)
	fmt.Println()

	//Day 32: Fibonacci Sequence
	fmt.Println("The 1st number in the Fibonacci Sequence: ",
		questions.FibonacciSequence(1))
	fmt.Println("The 17th number in the Fibonacci Sequence: ",
		questions.FibonacciSequence(17))
	fmt.Println("The 85th number in the Fibonacci Sequence: ",
		questions.FibonacciSequence(85))
	fmt.Println()

	//Day 33: Is valid BST
	fmt.Println("Is the Binary tree valid: ",
		tree.IsValid())
	fmt.Println()

	//Day 34: Bottom leftmost node
	fmt.Println("Bottom leftmost node's value: ",
		tree.BottomLeftNode())
	fmt.Println()
}
