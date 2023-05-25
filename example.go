package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"k8-go-cloud-micro/samples"
	"log"
	"math"
	_ "net/http/pprof"
	"sort"
	"strings"
	"unicode"
)

//Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	start, maxLen := 0, 0

	for i, char := range s {

		log.Printf("i: %d with character %c", i, char)
		if j, ok := seen[char]; ok && j >= start {
			log.Printf("j: %d, start: %d", j, start)
			start = j + 1
		}

		seen[char] = i
		currLen := i - start + 1

		if currLen > maxLen {
			maxLen = currLen
		}
	}

	log.Printf("seen %d with length: %d", seen, len(seen))
	return maxLen
}

//string s, reverse the order of the words.
func reverseOrder(s string) string {

	var temp string
	for i := len(s) - 1; i >= 0; i-- {
		temp += string(s[i])
	}
	return temp
}

func reverseString2nd(s string) string {
	alp := []rune(s)

	for i := 0; i < len(alp)/2; i++ {
		j := len(alp) - i - 1
		alp[i], alp[j] = alp[j], alp[i]
	}
	return string(alp)
}

func substringLargest(s string) (int, string) {
	seen := make(map[rune]int)

	temp := []rune(s)
	var longestSubString, currentSubString strings.Builder
	//var length, start int = 0
	for i := 0; i < len(temp); i++ {
		// condition
		if _, ok := seen[temp[i]]; ok {
			// reset the value
			currentSubString.Reset()
			seen = make(map[rune]int)
		}
		seen[temp[i]] = i
		currentSubString.WriteString(string(temp[i]))

		if len(longestSubString.String()) < len(currentSubString.String()) {
			longestSubString = currentSubString
		}
	}
	return len(longestSubString.String()), longestSubString.String()
}

//Given an array of integers, find the two elements that sum up to a given target.
func findPairs(a []int, k int) [][]int {
	seen := make(map[int]bool)
	var result [][]int

	for index, item := range a {
		el2 := k - item
		if _, ok := seen[el2]; ok {
			continue
		}
		for j := index + 1; j < len(a); j++ {
			if a[j] == el2 {
				result = append(result, []int{a[index], a[j]})
				seen[a[index]], seen[a[j]] = true, true
			}
		}
	}
	return result
}

//Given a string s, find the longest palindromic substring in s
func longestPalindrome(s string) string {
	// base condition
	if len(s) < 2 {
		return s
	}

	var output string
	runes := []rune(s)
	//var temp strings.Builder
	for i := range runes {
		for j := i + 1; j <= len(runes); j++ {
			if isPalindrome(string(runes[i:j])) {
				if len(output) < len(string(runes[i:j])) {
					output = string(runes[i:j])
				}
			}
		}
	}

	return output
}

func isPalindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[len(s)-i-1] {
			continue
		}
		return false
	}
	return true
}

//2nd Max element in array
func scndMaxElement(a []int) int {
	if len(a) <= 1 {
		return 0
	}
	var firstMax, secondMax int
	firstMax, secondMax = a[0], math.MinInt
	for i := 1; i < len(a); i++ {
		if firstMax > a[i] && secondMax < a[i] {
			secondMax = a[i]
		} else if a[i] > firstMax {
			secondMax = firstMax
			firstMax = a[i]
		}
	}
	return secondMax
}

//Replace a character in a string with and without using replace function
func replaceCharacter(s string, oldChar, newChar rune) string {
	runes := []rune(s)
	for index, item := range runes {
		if item == oldChar {
			runes[index] = newChar
		}
	}
	return string(runes)
}

// Replace a character in a string with and without using replace function
func replaceCharacterWithBuilder(s string, oldChar, newChar rune) string {
	var builder strings.Builder
	for _, ch := range s {
		if ch == oldChar {
			builder.WriteRune(newChar)
		} else {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}

//Problem: Given an array of integers, reverse the order of the elements in the array.
func reverseArray(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	return a
}

//Problem: Given an array of integers, find the first non-repeating element in the array.
func firstNonRepeatingElement(a []int) int {
	cal := make([][]int, 1000)
	for i := 0; i < len(a); i++ {
		cal[a[i]] = append(cal[a[i]], a[i])
	}

	for _, item := range a {
		if len(cal[item]) == 1 {
			return item
		}
	}
	return 0
}

func sortArray(array []int) []int {
	sort.Ints(array)
	return array
}

//Problem: Given two sorted arrays, merge them into a single sorted array.
func MergeTwoUnsortedArray(a, b []int) []int {
	i, j := 0, 0
	sortArray(a)
	sortArray(b)
	result := make([]int, 0)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else if a[i] > b[j] {
			result = append(result, b[j])
			j++
		} else if a[i] == b[j] {
			result = append(result, a[i], b[j])
			i++
			j++

		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}

	for j < len(b) {
		result = append(result, b[j])
	}
	return result
}

//Problem: Given an array of integers, remove duplicates from the array.
func removeDuplicates(array []int) []int {
	dedup := make(map[int]bool, len(array))
	result := make([]int, 0, len(array))
	for _, item := range array {
		if !dedup[item] {
			dedup[item] = true
			result = append(result, item)
		}
	}
	return result
}

//Problem: Given two arrays of integers, find the intersection of the arrays.
/*
Input:
a := []int{4, 9, 5}
b := []int{9, 4, 9, 8, 4}

Output:
[4, 9]
*/
func findIntersection(a, b []int) []int {
	isec := make(map[int]bool, len(a))
	result := make([]int, 0)
	for _, item := range a {
		isec[item] = true
	}

	for _, item := range b {
		if isec[item] {
			result = append(result, item)
			isec[item] = false
		}
	}
	return result
}

// rotten orange problem it will return the timeframe where all the oranges will be rotten
// 2 will be rotten 1 will be fresh and 0 means empty space
/*
2 1 1
1 1 0
0 1 1

*/
func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	freshOranges := 0
	queue := make([][2]int, 0)

	// Count the number of fresh oranges and add the rotten oranges to the queue
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				freshOranges++
			} else if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	minutes := 0
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for freshOranges > 0 && len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				r, c := curr[0]+dir[0], curr[1]+dir[1]

				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
					grid[r][c] = 2
					freshOranges--
					queue = append(queue, [2]int{r, c})
				}
			}
		}

		minutes++
	}

	if freshOranges > 0 {
		return -1
	}

	return minutes
}

//find-minimum-in-rotated-sorted-array -- done
//longest-increasing-subsequence -- working
//All 3 tree traversals
//Binary Search -- implement need to see again
//Write a sorting algorithm.
//Counting Coins

//find-minimum-in-rotated-sorted-array
func findMinRotatedSortedArray(array []int) int {
	if len(array) == 0 {
		return -1
	}
	min := array[0]
	for i := 0; i <= len(array)/2; i++ {
		if array[i] < min {
			min = array[i]
		} else if array[len(array)-i-1] < min {
			min = array[len(array)-i-1]
		}
	}
	return min
}

// binary search approach
func findMinRotatedSortedArrayBSA(array []int) int {
	if len(array) == 0 {
		return -1
	}
	left, right := 0, len(array)-1
	for left < right {
		mid := (left + right) / 2
		if array[mid] > array[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return array[0]
}

// longest-increasing-subsequence (dynamic programming (dp)) O(n^2).
//Input: [10,9,2,5,3,7,101,18]
//Output: 4
func longestSubSeq(array []int) int {
	if len(array) == 0 {
		return 0
	}
	dp := make([]int, len(array))
	dp[0] = 1
	maxLen := 1
	for i := 1; i < len(array); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			//log.Println("Details:", array[j], array[i], dp[i], dp[j]+1)
			if array[j] < array[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

func longestSubSequence(array []int) int {
	maxLen := 1
	dp := make([]int, len(array))
	dp[0] = 1
	for i := 1; i < len(array); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if array[j] < array[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	log.Println("DP", dp)
	return maxLen
}

// counting coins
//Coins: [1, 2, 5], amount: 11
//Output: 3

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// initialize the dp array with maximum possible value
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	// loop through each coin and calculate the minimum number of coins needed for each amount
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	// check if it's possible to make the given amount with the given coins
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// helper function to calculate the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// with dp used
func countCoinsdp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

//find the count of same number with there count ?
//for eg: 1 - 3 times
//2 - 5 times
//arr := {1,1,1,2,2,2,2,2} in golang

func countArray(array []int) map[int]int {
	result := make(map[int]int, len(array))
	for i := 0; i < len(array); i++ {
		result[array[i]] += 1
	}
	return result
}

//Odd pair intersection
//1,3,5,7
func oddPairIntersection(arr, arr2 []int) []int {
	temp := make(map[int]bool, len(arr))
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			temp[arr[i]] = true
		}
	}

	for i := 0; i < len(arr2); i++ {
		if temp[arr2[i]] {
			temp[arr2[i]] = false
			result = append(result, arr2[i])
		}
	}
	return result

}

//minimum element in a rotated sorted array using golang, we can use a binary search approach
func findMinInRotatedSorted(nums []int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

//Balance parenthesis
//time and space complexity O(n)
func balanceParenthesis(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) == 0 || m[c] != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

//Given two list and and a X, find out such a pair which sum equal to X.
func findPairWithSum(a, b []int, target int) (int, int, bool) {
	set := make(map[int]bool)
	for _, num := range a {
		set[num] = true
	}
	for _, num := range b {
		complement := target - num
		if set[complement] {
			return complement, num, true
		}
	}
	return 0, 0, false
}

// find largest duplicacy element in an array ex:=a=[1,1,1,1,2,2,3,3,3,3]
//output 3 bcz 3 & 1 are having equal freq but max is 3 so ans is 3
func largestDuplicacy(arr []int) int {
	temp := make(map[int]int, len(arr))
	for _, item := range arr {
		temp[item]++
	}

	maxNum, maxVal := 0, 0
	for key, value := range temp {
		if key > maxNum {
			if value >= maxVal {
				maxNum = key
				maxVal = value
			}
		}
		if value > maxVal {
			maxNum = key
			maxVal = value
		}
	}
	return maxNum
}

/*
Q. Find the greatest element in a string.
 E.g "abc345yue674iut789"
 Output: 789
*/
func findGreatestElementInString(s string) int {
	max := 0
	temp := ""
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			temp += string(ch)
		} else if temp != "" {
			num := parseInt(temp)
			if num > max {
				max = num
			}
			temp = ""
		}
	}

	// Check if the last character was a digit
	if temp != "" {
		num := parseInt(temp)
		if num > max {
			max = num
		}
	}
	return max
}

// parse int
func parseInt(s string) int {
	n := 0
	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}
	return n
}

// Given a list of 0's and 1's and find the first occurence of 1 in best time complexity.
// we can use binary search
// time complexity to O(log n)
func firstOccurrence(arr []int) int {
	low := 0
	high := len(arr) - 1
	result := -1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == 1 {
			result = mid
			high = mid - 1
		} else if arr[mid] < 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func lruCache() {
	// Test case 1
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	//	[null, null, null, 1, null, -1, null, -1, 3, 4]
	fmt.Println("Test case 1")
	obj := samples.Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

	// Test case 2
	// 	["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]
	fmt.Println("Test case 2")
	obj = samples.Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}

// Given a list of stock prices, find maximum profit. You can buy one day and sell on another.
func getMaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func calculateHCF(a, b int) int {
	if b == 0 {
		return a
	}

	return calculateHCF(b, a%b)
}

func fibonacci(n int) []int {
	sequence := []int{0, 1}

	for i := 2; i < n; i++ {
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return sequence
}

func findUniqueElements(list1, list2 []int) []int {
	uniqueList := make([]int, 0)
	seen := make(map[int]bool)

	// Iterate over list1 and add unique elements to the uniqueList
	for _, num := range list1 {
		if !seen[num] {
			uniqueList = append(uniqueList, num)
			seen[num] = true
		}
	}

	// Iterate over list2 and add unique elements to the uniqueList
	for _, num := range list2 {
		if !seen[num] {
			uniqueList = append(uniqueList, num)
			seen[num] = true
		}
	}

	return uniqueList
}

// shorten url
func shortenURL(longURL string) string {
	hash := md5.Sum([]byte(longURL))
	shortURL := hex.EncodeToString(hash[:])

	// Trim the hash to desired length
	shortURL = shortURL[:8]

	return shortURL
}

// find the string without using the inbuild function
func stringLength(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func main() {

	s := "abcdeefababahdabcdefghijklabcdefnopqxzyzyxbcapqrstuvw"
	var array []int
	array = append(array, 1, 10, 9, 2, 5, 3, 7, 101, 18, 2, 1)
	grid := [][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(reverseOrder(s))
	fmt.Println(reverseString2nd(s))
	fmt.Println(substringLargest(s))
	fmt.Println(findPairs(array, 9))
	fmt.Println(longestPalindrome(s))
	fmt.Println("Second Max Element from array:", scndMaxElement(array))
	fmt.Println("Replace Character from String a convert to #:", replaceCharacter(s, 'a', '#'))
	fmt.Println("Reverse Array:", reverseArray(array))
	fmt.Println("First Non repeating char:", firstNonRepeatingElement(array))

	fmt.Println("Minites of rotten oranges:", orangesRotting(grid))
	//Given a list of stock prices, find maximum profi
	fmt.Println("Maximum Profit:", getMaxProfit([]int{2, 4, 1}))

	fmt.Println("Merge Unsorted array:", MergeTwoUnsortedArray(array, []int{1, 1, 2}))
	fmt.Println("Remove Duplicates:", removeDuplicates(array))
	fmt.Println("Intersection:", findIntersection(array, []int{1, 3, 4, 1, 2, 1}))
	fmt.Println("Find Min in Array:", findMinRotatedSortedArray(array))
	fmt.Println("Find Min in Array binary search approach ", findMinRotatedSortedArrayBSA(array))
	fmt.Println("Longest Sub Sequence:", longestSubSeq(array))
	fmt.Println("Longest Sub Sequence other:", longestSubSequence(array))
	fmt.Println("count coins:", coinChange([]int{1, 2, 5}, 11))
	fmt.Println("count coind:", countCoinsdp([]int{1, 2, 5}, 11))
	//fmt.Println("Test:", test140(array))
	fmt.Println("Count the Numbers:", countArray(array))
	fmt.Println("Odd Pair Intersection:", oddPairIntersection(array, array))
	//lruCache()
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(samples.MinimumTotal(triangle))
	s1 := "(1+(2*3)+((8)/4))+1"
	fmt.Printf("The maximum depth of parentheses in '%s' is %d\n", s1, samples.MaxDepthAdvances(s1))

	a := []int{1, 3, 5, 7}
	elem := 4
	index := samples.BinarySearchIndex(a, elem)
	fmt.Printf("Element %d should be inserted at index %d in %v\n", elem, index, a)

	fmt.Println("minimum element in a rotated sorted array:", findMinInRotatedSorted(array))

	p := "()[]{}"
	fmt.Println(p, balanceParenthesis(p)) // true

	// find pair with sum
	a = []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	x := 10
	num1, num2, ok := findPairWithSum(a, b, x)
	if ok {
		fmt.Printf("Found pair: %d, %d\n", num1, num2)
	} else {
		fmt.Println("No pair found")
	}

	// Given a list of 0's and 1's and find the first occurence of 1
	fmt.Println("first occurrence", firstOccurrence(array))
	//
	fmt.Println("Greatest Element from String:", findGreatestElementInString("abc345yue674iut789"))

	// Server init
	//ServerInit()

	fmt.Println("Largest Duplicate Number:", largestDuplicacy([]int{1, 1, 1, 1, 5, 5, 5, 5, 2, 2, 3, 3, 3, 3}))

	fmt.Println("Prime Number:", isPrime(17))
	fmt.Println("HCF:", calculateHCF(36, 48))
	fmt.Println("fibonacci:", fibonacci(10))
	fmt.Println("UniqueElements:", findUniqueElements([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}))

	fmt.Println("Short URL:", shortenURL("https://www.abc.com"))

}
