package test

import (
	"github.com/mrbarge/aoc2024-golang/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("File helpers", func() {

	Context("ReadLines", func() {
		It("Behaves correctly", func() {
			in := `abcd
efgh
ijkl`
			ir := strings.NewReader(in)
			out := []string{"abcd", "efgh", "ijkl"}
			res, err := helper.ReadLines(ir, false)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
		It("Ignores empty lines", func() {
			in := `abcd

efgh`
			ir := strings.NewReader(in)
			out := []string{"abcd", "efgh"}
			res, err := helper.ReadLines(ir, true)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})

	})

	Context("ReadLinesAsInt", func() {
		It("Behaves correctly", func() {
			in := `123
-456
0
3322351`
			ir := strings.NewReader(in)
			out := []int{123, -456, 0, 3322351}
			res, err := helper.ReadLinesAsInt(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
		It("Ignores empty lines", func() {
			in := `123

0`
			ir := strings.NewReader(in)
			out := []int{123, 0}
			res, err := helper.ReadLinesAsInt(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
	})
	Context("ReadLinesAsIntArray", func() {
		It("Behaves correctly", func() {
			in := `123
456
7892
`
			ir := strings.NewReader(in)
			out := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9, 2}}
			res, err := helper.ReadLinesAsIntArray(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
	})

	Context("with a CSV file-like input", func() {
		It("parses data with a comma delimiter correctly", func() {
			// Test input data (simulating a file input with a string)
			input := "10,20,30\n40,50,60\n70,80,90"

			// Use strings.NewReader to simulate file input
			data, err := helper.ReadLinesAsCSVIntArray(strings.NewReader(input), ",")

			// Assertions
			Expect(err).To(BeNil()) // Ensure there is no error
			Expect(data).To(Equal([][]int{
				{10, 20, 30},
				{40, 50, 60},
				{70, 80, 90},
			})) // Check if data is correctly parsed
		})

		It("parses data with a semicolon delimiter correctly", func() {
			// Test input data (simulating a file input with a string)
			input := "10;20;30\n40;50;60\n70;80;90"

			// Use strings.NewReader to simulate file input
			data, err := helper.ReadLinesAsCSVIntArray(strings.NewReader(input), ";")

			// Assertions
			Expect(err).To(BeNil()) // Ensure there is no error
			Expect(data).To(Equal([][]int{
				{10, 20, 30},
				{40, 50, 60},
				{70, 80, 90},
			})) // Check if data is correctly parsed
		})

		It("returns an error when there is a malformed line", func() {
			// Test input data with an invalid line
			input := "10,20,30\ninvalid,entry\n40,50,60"

			// Use strings.NewReader to simulate file input
			data, err := helper.ReadLinesAsCSVIntArray(strings.NewReader(input), ";")

			// Assertions
			Expect(err).ToNot(BeNil()) // There should be an error
			Expect(data).To(BeNil())   // No valid data should be parsed
		})
	})
})
