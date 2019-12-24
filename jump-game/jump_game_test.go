package jumpgame

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJumpGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jump Game Test Suite")
}

var _ = Describe("canJump", func() {
	Context("With valid input", func() {
		It("should determine that you are able to reach the last index", func() {
			input := []int{2, 3, 1, 1, 4}
			Expect(canJump(input)).To(Equal(true))
		})
	})

	Context("With invalid input", func() {
		It("should determine that you are not able to reach the last index", func() {
			input := []int{3, 2, 1, 0, 4}
			Expect(canJump(input)).To(Equal(false))
		})
	})
})
