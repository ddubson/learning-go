package adder_test

import (
	"github.com/ddubson/learning-go/09-testing/adder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
BDD Test with Ginkgo
To execute, run `go test 09-testing`
*/
var _ = Describe("Adder Suite", func() {
	Describe("Add", func() {
		It("adds two numbers together", func() {
			Expect(adder.Add(2, 2)).To(Equal(4))
		})
	})
})
