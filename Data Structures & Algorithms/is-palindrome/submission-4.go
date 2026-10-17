// there's unicode function

func isPalindrome(s string) bool {
	left := 0
	right := len(s)-1

	// two pointers function, check from left, and check from right
	for left < right{
		for left < right && !isAlphaNumeric(rune(s[left])) {
            left++ // damn traversal, but stil O(n)
        }
        for right > left && !isAlphaNumeric(rune(s[right])) {
            right-- // damn traversal O(n)
        }
        if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
            return false // check if mismatch
        }
        left++
        right--
    }
    return true
}

func isAlphaNumeric(char rune) bool {
    return unicode.IsLetter(char) || unicode.IsDigit(char)
}