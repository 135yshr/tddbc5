package tddbc5_test

import (
	. "github.com/135yshr/tddbc5"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("格子点", func() {
	Describe("格子点を文字列で表す", func() {
		Context("座標Xが4、座標Yが7の格子点を作成する", func() {
			It("文字列表記が(4,7)になること", func() {
				sut := GridPoint{4, 7}
				Expect(sut.Notation()).To(Equal("(4,7)"))
			})
		})
	})
})
