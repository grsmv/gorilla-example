package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/gorilla/mux"
	"github.com/drewolson/testflight"
	"github.com/grsmv/gorilla-example/app"
)

var r = mux.NewRouter()

var _ = Describe("SampleController", func() {

	BeforeSuite(func(){
		app.RegisterRoutes(r)
	})

	Context("GET /", func(){
		It("should be OK", func(){
			testflight.WithServer(r, func(r *testflight.Requester){
				response := r.Get("/")
				Expect(response.StatusCode).To(Equal(200))
			})
		})
	})

	Context("GET /books", func(){
		It("Should be OK too", func(){
			testflight.WithServer(r, func(r *testflight.Requester){
				response := r.Get("/books")
				Expect(response.StatusCode).To(Equal(200))
			})
		})
	})
})
