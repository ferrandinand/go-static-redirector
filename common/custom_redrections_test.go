package common

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func RedirectionsTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CustomRedirectionsTest")
}

/*
var _ = Describe("CustomRedirectionsTest", func() {
	Context("JSON File is loaded properly", func() {
		It("Redirects", func() {

			var jsonEquivalencies = []struct {
				Path string
				URL  string
			}{
				{Path: "/health", URL: "http://www.flywire.com/healthcare"},
				{Path: "/text", URL: "/music"},
			}
			for _, url := range romanequivalences {
				Expect(Convertarabictoroman(arabicromanmap.arabic)).Should(BeEquivalentTo(arabicromanmap.roman))
			}			Expect(CustomRedirections()).To(Equal(0.0))
		})
	})
})
*/
