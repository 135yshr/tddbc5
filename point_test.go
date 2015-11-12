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
			It("yが1になること", func() {
				p := Point{X: 1, Y: 1}
				Expect(p.Y).To(Equal(1))
			})
		})
		Context("xに2 yに3を渡したとき", func() {
			It("xが2になること", func() {
				p := Point{X: 2, Y: 3}
				Expect(p.X).To(Equal(2))
			})
			It("yが3になること", func() {
				p := Point{X: 2, Y: 3}
				Expect(p.Y).To(Equal(3))
			})
		})
	})
})
