package plusone

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlusOne(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plus One Test Suite")
}

var _ = Describe("plusOne", func() {
	Context("Given [1, 2, 3]", func() {
		It("should increment a given slice", func() {
			input := []int{1, 2, 3}
			expected := []int{1, 2, 4}
			Expect(plusOne(input)).To(Equal(expected))
		})
	})
	Context("Given [4, 3, 2, 1]", func() {
		It("should increment a given slice", func() {
			input := []int{4, 3, 2, 1}
			expected := []int{4, 3, 2, 2}
			Expect(plusOne(input)).To(Equal(expected))
		})
	})
	Context("Given [9]", func() {
		It("should properly prepend one", func() {
			input := []int{9}
			expected := []int{1, 0}
			Expect(plusOne(input)).To(Equal(expected))
		})
	})
	Context("Given [1, 9]", func() {
		It("should properly increment the slice", func() {
			input := []int{1, 9}
			expected := []int{2, 0}
			Expect(plusOne(input)).To(Equal(expected))
		})
	})
})
