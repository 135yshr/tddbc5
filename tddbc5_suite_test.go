package tddbc5_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTddbc5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tddbc5 Suite")
}
