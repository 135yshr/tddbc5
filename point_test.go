package tddbc5_test

import (
	. "github.com/135yshr/tddbc5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("平面座標上の点", func() {
	Describe("座業を作成する", func() {
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
	Describe("座業が隣り合っていることを確認する", func() {
		Context("座標A(1,1)と座業B(1,2)が存在するとき", func() {
			It("座標Aと座標Bが隣り合っていること", func() {
				Expect(a.IsNeighbor(b)).To(BeTrue())
			})
		})
	})
})
