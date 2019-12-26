package validpalindrome

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidPalindrome(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Valid Palindrome Test Suite")
}

var _ = Describe("isPalindrome", func() {
	Context("Given a valid string: 'A man, a plan, a canal: Panama'", func() {
		It("Should determine that it is a palindrome", func() {
			input := "A man, a plan, a canal: Panama"
			Expect(isPalindrome(input)).To(Equal(true))
		})
	})
	Context("Given an invalid string: 'race a car'", func() {
		It("Should determine that it is not a palindrome", func() {
			input := "race a car"
			Expect(isPalindrome(input)).To(Equal(false))
		})
	})
})
