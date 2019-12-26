package reversebits

import (
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReverseBits(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reverse Bits Test Suite")
}

var _ = Describe("reverseBits", func() {
	Context("Given a 32 bits unsigned integer", func() {
		It("Should reverse the integer", func() {
			in, _ := strconv.ParseUint("00000010100101000001111010011100", 2, 32)
			out, _ := strconv.ParseUint("00111001011110000010100101000000", 2, 32)
			input := uint32(in)
			expected := uint32(out)
			Expect(reverseBits(input)).To(Equal(expected))
		})
	})
})
