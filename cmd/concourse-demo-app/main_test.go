package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/concourse-demo-app/cmd/concourse-demo-app"
)

var _ = Describe("main", func() {
	Describe("given IndexController() function", func() {
		Context("when handling an http request", func() {
			var responseWriter *httptest.ResponseRecorder
			var request *http.Request

			BeforeEach(func() {
				responseWriter = httptest.NewRecorder()
				IndexController(responseWriter, request)
			})
			It("then it should return a nice welcome message asldkjfasdf", func() {
				body, _ := ioutil.ReadAll(responseWriter.Body)
				Ω(string(body)).Should(Equal(WelcomeMessage))
			})
		})
	})
})
