package source_test

import (
	"github.com/ddubson/learning-go/source"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello", func() {
	It("says hello", func() {
		Expect(source.Hello()).To(Equal("Hello, World!"))
	})
})
