package main

import (
	"testing"

	"github.com/ferrandinand/gocheck"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err    error
	Resp   string
	Status int
}

func TestGocheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go check Suite")
}

var _ = Describe("Go check Suite", func() {
	Context("Google get correct output", func() {

		It("Returns 200", func() {
			endpoints := []string{"http://localhost:8080/text", "http://localhost:8080/music"}
			result, _ := gocheck.BarrierStatusCode(endpoints...)
			Expect(result[0]).To(Equal(200))
		})

		It("Returns 404", func() {
			endpoints := []string{"http://localhost:8080/fail"}
			result, _ := gocheck.BarrierStatusCode(endpoints...)
			Expect(result[0]).To(Equal(404))
		})

	})
})
