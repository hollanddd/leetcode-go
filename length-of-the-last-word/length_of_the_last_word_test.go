package lengthlastword

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLengthLastWord(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Length of the Last Word Test Suite")
}

var _ = Describe("lengthOfLastWord", func() {
	Context("Given 'Hello World'", func() {
		It("should return the length of the last word", func() {
			input := "Hello World"
			Expect(lengthOfLastWord(input)).To(Equal(5))
		})
	})
	Context("Given 'a '", func() {
		It("should exclude single character prhase", func() {
			input := "a "
			Expect(lengthOfLastWord(input)).To(Equal(1))
		})
	})
	Context("Given ''", func() {
		It("should exclude single character prhase", func() {
			input := ""
			Expect(lengthOfLastWord(input)).To(Equal(0))
		})
	})
	Context("Given ' '", func() {
		It("should exclude single character prhase", func() {
			input := " "
			Expect(lengthOfLastWord(input)).To(Equal(0))
		})
	})
	Context("Given '     '", func() {
		It("should exclude single character prhase", func() {
			input := "     "
			Expect(lengthOfLastWord(input)).To(Equal(0))
		})
	})
	Context("Given '  day'", func() {
		It("should exclude single character prhase", func() {
			input := "  day"
			Expect(lengthOfLastWord(input)).To(Equal(3))
		})
	})
})
