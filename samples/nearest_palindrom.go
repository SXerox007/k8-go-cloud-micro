package samples

//Get nearest palindrome of the given number

import (
	"fmt"
	"strconv"
)

func mainNearestPalindrom() {
	number := 12345
	nearestPalindrome := findNearestPalindrome(number)
	fmt.Println("Nearest Palindrome:", nearestPalindrome)
}

func findNearestPalindrome(number int) int {
	// Check if the number itself is a palindrome
	if isPalindrome(number) {
		return number
	}

	// Find the nearest smaller palindrome
	smallerPalindrome := findSmallerPalindrome(number)

	// Find the nearest larger palindrome
	largerPalindrome := findLargerPalindrome(number)

	// Compare the differences and return the nearest palindrome
	diffSmaller := number - smallerPalindrome
	diffLarger := largerPalindrome - number

	if diffSmaller < diffLarger {
		return smallerPalindrome
	}

	return largerPalindrome
}

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func findSmallerPalindrome(number int) int {
	for i := number - 1; i >= 0; i-- {
		if isPalindrome(i) {
			return i
		}
	}

	return -1 // No smaller palindrome found
}

func findLargerPalindrome(number int) int {
	for i := number + 1; ; i++ {
		if isPalindrome(i) {
			return i
		}
	}

	return -1 // No larger palindrome found
}
