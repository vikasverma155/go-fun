package cracking_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/vikasverma155/go-fun/learn/algos/hackerrank/cracking"
)

var _ = Describe("Small", func() {
	It("should generate fibonacci", func() {
		Expect(Fibonacci(34)).To(Equal(5702887))
	})
})
