package tddbc5_test

import (
	. "github.com/135yshr/tddbc5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Point", func() {
	Describe("平面座標上の点", func() {
		var p Point
		Context("xに1 yに1を渡したとき", func() {
			BeforeEach(func() {
				p = Point{X: 1, Y: 1}
			})
			It("xが1になること", func() {
				Expect(p.X).To(Equal(1))
			})
			It("yが1になること", func() {
				Expect(p.Y).To(Equal(1))
			})
		})
		Context("xに2 yに3を渡したとき", func() {
			BeforeEach(func() {
				p = Point{X: 2, Y: 3}
			})
			It("xが2になること", func() {
				Expect(p.X).To(Equal(2))
			})
			It("yが3になること", func() {
				Expect(p.Y).To(Equal(3))
			})
		})
	})
})
