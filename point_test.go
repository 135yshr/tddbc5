package tddbc5_test

import (
	. "github.com/135yshr/tddbc5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("格子点", func() {
	Describe("格子点を文字列で表す", func() {
		Context("座標Xが4、座標Yが7の格子点を作成したとき", func() {
			It("文字列表記が(4,7)になること", func() {
				sut := GridPoint{4, 7}
				Expect(sut.Notation()).To(Equal("(4,7)"))
			})
		})
		Context("座標Xが-2、座標Yが3の格子点を作成したとき", func() {
			It("文字列表記が(-2,3)になること", func() {
				sut := GridPoint{-2, 3}
				Expect(sut.Notation()).To(Equal("(-2,3)"))
			})
		})
	})
	Describe("２つの格子点が同じ座標を持つか判定する", func() {
		Context("格子点oneが(4,7)、格子点otherが(4,7)のとき", func() {
			It("格子点oneと格子点otherが同じを持っていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{4, 7}
				Expect(sut.SameCoordinatesWith(other)).To(BeTrue())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(10, 8)のとき", func() {
			It("格子点oneと格子点otherが異なる座標を持っていないこと", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{10, 8}
				Expect(sut.SameCoordinatesWith(other)).To(BeFalse())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(4, 9)のとき", func() {
			It("格子点oneと格子点otherが異なる座標を持っていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{4, 9}
				Expect(sut.SameCoordinatesWith(other)).To(BeFalse())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(10, 7)のとき", func() {
			It("格子点oneと格子点otherが異なる座標を持っていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{10, 7}
				Expect(sut.SameCoordinatesWith(other)).To(BeFalse())
			})
		})
	})
	Describe("２つの格子点が隣り合っているか判定する", func() {
		Context("格子点oneが(4,7)、格子点otherが(4,8)のとき", func() {
			It("格子点oneと格子点otherは隣なっていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{4, 8}
				Expect(sut.NeighborOf(other)).To(BeTrue())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(4,6)のとき", func() {
			It("格子点oneと格子点otherは隣なっていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{4, 6}
				Expect(sut.NeighborOf(other)).To(BeTrue())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(3,7)のとき", func() {
			It("格子点oneと格子点otherは隣なっていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{3, 7}
				Expect(sut.NeighborOf(other)).To(BeTrue())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(5,7)のとき", func() {
			It("格子点oneと格子点otherは隣なっていること", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{5, 7}
				Expect(sut.NeighborOf(other)).To(BeTrue())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(4,7)のとき", func() {
			It("格子点oneと格子点otherは隣なっていないこと", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{4, 7}
				Expect(sut.NeighborOf(other)).To(BeFalse())
			})
		})
		Context("格子点oneが(4,7)、格子点otherが(3,8)のとき", func() {
			It("格子点oneと格子点otherは隣なっていないこと", func() {
				sut := GridPoint{4, 7}
				other := GridPoint{3, 8}
				Expect(sut.NeighborOf(other)).To(BeFalse())
			})
		})
	})
})
