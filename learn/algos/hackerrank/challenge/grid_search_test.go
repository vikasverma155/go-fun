package challenge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bufio"

	"fmt"

	"github.com/vikasverma155/go-fun/learn/algos/hackerrank/challenge"
	"github.com/vikasverma155/go-fun/util"
)

var _ = Describe("GridSearch", func() {
	It("should work case 1", func() {
		scanner := util.NewStringScanner(input1)
		c := util.ReadInt(scanner)
		output := []bool{true, false}
		for i := 0; i < c; i++ {
			grid, search := readSet(scanner)
			Expect(challenge.GridSearch(grid, search)).To(Equal(output[i]), fmt.Sprintf("Failed Case %v", i+1))
		}

	})
})

func readSet(scanner *bufio.Scanner) (grid, search []string) {
	nm := util.ReadInts(scanner, 2)
	grid = util.ReadStrings(scanner, nm[0])
	nm = util.ReadInts(scanner, 2)
	search = util.ReadStrings(scanner, nm[0])
	return
}

var (
	input1 = `2
10 10
7283455864
6731158619
8988242643
3830589324
2229505813
5633845374
6473530293
7053106601
0834282956
4607924137
3 4
9505
3845
3530
15 15
400453592126560
114213133098692
474386082879648
522356951189169
887109450487496
252802633388782
502771484966748
075975207693780
511799789562806
404007454272504
549043809916080
962410809534811
445893523733475
768705303214174
650629270887160
2 2
99
99`
)
