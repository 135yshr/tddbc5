package tddbc5_test

import (
	. "github.com/135yshr/tddbc5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Point", func() {
	Describe("平面座標上の点", func() {
		Context("xに1 yに1を渡したとき", func() {
			It("xが1になること", func() {
				p := Point{X: 1, Y: 1}
				Expect(p.X).To(Equal(1))
			})
		})
	})
})
