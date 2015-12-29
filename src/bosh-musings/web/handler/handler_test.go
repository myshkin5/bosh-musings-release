package handler_test

import (
	"bosh-musings/web/handler"

	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	It("returns the expected response", func() {
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/", nil)
		Expect(err).NotTo(HaveOccurred())

		handler.ServeHTTP(recorder, request)

		Expect(recorder.Body.String()).To(Equal("hello, world!\n"))
	})
})
