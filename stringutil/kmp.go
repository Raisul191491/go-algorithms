package stringutil

// KMPSearch performs the Knuth-Morris-Pratt (KMP) algorithm to check if the pattern exists in the text.
func KMPSearch(pattern string, text string) bool {
	patternLength := len(pattern)
	textLength := len(text)

	// Early return if the pattern is longer than the text
	if patternLength > textLength {
		return false
	}

	// Compute the Longest Prefix Suffix (LPS) array for the pattern
	lps := computeLPSArray(pattern)

	patternIndex, textIndex := 0, 0

	for textIndex < textLength {
		if pattern[patternIndex] == text[textIndex] {
			patternIndex++
			textIndex++
		}

		// If the entire pattern is found in the text
		if patternIndex == patternLength {
			return true
		}

		// Mismatch after patternIndex matches
		if textIndex < textLength && pattern[patternIndex] != text[textIndex] {
			if patternIndex != 0 {
				patternIndex = lps[patternIndex-1]
			} else {
				textIndex++
			}
		}
	}

	return false
}

// computeLPSArray calculates the Longest Prefix Suffix (LPS) array for a given pattern.
// The LPS array is used to skip characters during the matching process in the KMP algorithm.
func computeLPSArray(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0 // Length of the previous longest prefix suffix
	i := 1      // Start from the second character

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}
