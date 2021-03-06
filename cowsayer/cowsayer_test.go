package cowsayer_test

import (
	c "github.com/clijockey/cowsay/cowsayer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cowsayer package", func() {
	Describe("Simplesay()", func() {
		It("should export Simplesay", func() {
			Ω(func() { c.Simplesay([]string{}) }).Should(Not(Panic()))
		})

		// Context("when cowsay with another string is passed in", func() {
		// 	var (
		// 		params []string
		// 	)
		// 	BeforeEach(func() {
		// 		params = []string{"cowsay"}
		// 	})

		// 	Context("when 'rob' is also passed in", func() {
		// 		BeforeEach(func() {
		// 			params = append(params, "rob")
		// 		})

		// 		It("returns 'oh rob'", func() {
		// 			Ω(c.Simplesay(params)).Should(Equal("oh rob"))
		// 		})
		// 	})

		// 	Context("when'stephan' is also passed in", func() {

		// 		BeforeEach(func() {
		// 			params = []string{"cowsay", "stephan"}
		// 		})

		// 		It("returns 'oh stephan'", func() {
		// 			Ω(c.Simplesay(params)).Should(Equal("oh stephan"))
		// 		})
		// 	})
		// })

		// Context("when cowsay with multiple string is passed in", func() {
		// 	var (
		// 		params []string
		// 	)
		// 	BeforeEach(func() {
		// 		params = []string{"cowsay"}
		// 	})

		// 	Context("when 'rob' and 'stephan' is also passed in", func() {
		// 		BeforeEach(func() {
		// 			params = append(params, "rob", "stephan")
		// 		})

		// 		It("returns 'oh rob stephan'", func() {
		// 			Ω(c.Simplesay(params)).Should(Equal("oh rob stephan"))
		// 		})
		// 	})

		// })

		// Same as above however in a better format
		var saytest = []struct {
			in  []string
			out string
		}{
			{[]string{"cowsay"}, " ________\n< oh hey >\n --------\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||"},
			{[]string{"cowsay", "rob"}, " _____\n< rob >\n -----\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||"},
			{[]string{"cowsay", "super", "duppa"}, " _____________\n< super duppa >\n -------------\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||"},
		}

		for _, tt := range saytest {
			It("works", func() {
				Ω(c.Simplesay(tt.in)).Should(Equal(tt.out))
			})
		}
	})
})
