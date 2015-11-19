package tddbc5_test

import (
	. "github.com/135yshr/tddbc5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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
	DescribeTable("２つの格子点が同じ座標を持つか判定する",
		func(a GridPoint, expected bool) {
			sut := GridPoint{4, 7}
			Expect(sut.SameCoordinatesWith(a)).To(Equal(expected))
		},
		Entry("対象と同じ座標を持つ格子点", GridPoint{4, 7}, true),
		Entry("全く異なる座標を持つ格子点", GridPoint{10, 8}, false),
		Entry("Y座標が１つ違う格子点", GridPoint{4, 8}, false),
		Entry("X座標が１つ違う格子点", GridPoint{3, 7}, false),
	)
	DescribeTable("２つの格子点が隣り合っていること",
		func(a GridPoint) {
			sut := GridPoint{4, 7}
			Expect(sut.NeighborOf(a)).To(BeTrue())
		},
		Entry("(4,7) (3,7)", GridPoint{3, 7}),
		Entry("(4,7) (5,7)", GridPoint{5, 7}),
		Entry("(4,7) (4,6)", GridPoint{4, 6}),
		Entry("(4,7) (4,8)", GridPoint{4, 8}),
	)
	DescribeTable("２つの格子点が隣り合っていないこと",
		func(sut, other GridPoint) {
			Expect(sut.NeighborOf(other)).To(BeFalse())
		},
		Entry("(4,7) (3,6)", GridPoint{4, 7}, GridPoint{3, 6}),
		Entry("(4,7) (3,8)", GridPoint{4, 7}, GridPoint{3, 8}),
		Entry("(4,7) (5,6)", GridPoint{4, 7}, GridPoint{5, 6}),
		Entry("(4,7) (5,8)", GridPoint{4, 7}, GridPoint{5, 8}),
		Entry("(4,7) (4,7)", GridPoint{4, 7}, GridPoint{4, 7}),
	)
})
