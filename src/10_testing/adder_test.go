package main_test

import (
	. "10_testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
	BDD Test with Ginkgo
	To execute, run `go test 10_testing`
 */
var _ = Describe("Poem Test Suite", func() {
	Describe("Add", func() {
		It("adds two numbers together", func() {
			Expect(Add(2, 2)).To(Equal(4))
		})
	})
})
