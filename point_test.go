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
		Context("格子点Aが(4,7)、格子点Bが(4,7)のとき", func() {
			It("格子点Aと格子点Bが同じであること", func() {
				a := GridPoint{4, 7}
				b := GridPoint{4, 7}
				Expect(a.SameCoordinatesWith(b)).To(BeTrue())
			})
		})
	})
})
