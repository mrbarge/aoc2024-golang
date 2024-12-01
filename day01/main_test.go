package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("day01", func() {

	Context("When sorting the data", func() {
		It("behaves successfully with unique values", func() {
			testdata := [][]int{
				{3, 1},
				{5, 2},
				{1, 5},
				{2, 3},
			}
			l, r := getSortedSides(testdata)
			Expect(l).To(Equal([]int{1, 2, 3, 5}))
			Expect(r).To(Equal([]int{1, 2, 3, 5}))
		})

		It("behaves successfully with duplicate values", func() {
			testdata := [][]int{
				{3, 1},
				{5, 2},
				{1, 5},
				{2, 3},
				{3, 2},
			}
			l, r := getSortedSides(testdata)
			Expect(l).To(Equal([]int{1, 2, 3, 3, 5}))
			Expect(r).To(Equal([]int{1, 2, 2, 3, 5}))
		})

		It("behaves successfully when checking similarity", func() {
			s := getSimilarity(
				[]int{3, 5, 1, 2, 6},
				[]int{1, 2, 5, 2, 2})
			Expect(s[1]).To(Equal(1))
			Expect(s[2]).To(Equal(3))
			Expect(s[3]).To(Equal(0))
			Expect(s[5]).To(Equal(1))
			Expect(s[6]).To(Equal(0))
		})

	})
})
