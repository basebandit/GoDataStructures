package palindrome

func isPalindrome(n int) bool {
	var rev, num int
	num = n

	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev == num
}
