package main

import (
	"github.com/mrbarge/aoc2024-golang/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Base", func() {

	Context("Perimeter", func() {
		It("behaves successfully with unique values", func() {
			// shapes are from the problem description
			dshape := Region{
				Coords: map[helper.Coord]struct{}{
					{0, 0}: {},
				},
			}
			cshape := Region{
				Coords: map[helper.Coord]struct{}{
					{0, 0}: {},
					{0, 1}: {},
					{1, 1}: {},
					{1, 2}: {},
				},
			}
			eshape := Region{
				Coords: map[helper.Coord]struct{}{
					{3, 3}: {},
					{4, 3}: {},
					{5, 3}: {},
				},
			}
			Expect(dshape.Perimeter()).To(Equal(4))
			Expect(cshape.Perimeter()).To(Equal(10))
			Expect(eshape.Perimeter()).To(Equal(8))

			Expect(dshape.Area()).To(Equal(1))
			Expect(cshape.Area()).To(Equal(4))
			Expect(eshape.Area()).To(Equal(3))
		})

	})
})
