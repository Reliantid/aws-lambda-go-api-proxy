package chiadapter_test

import (
	"context"
	"log"
	"net/http"

	"github.com/Reliantid/aws-lambda-go-api-proxy/chi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-chi/chi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChiLambda tests", func() {
	Context("Simple ping request", func() {
		It("Proxies the event correctly", func() {
			ctx := context.Background()
			log.Println("Starting test")

			r := chi.NewRouter()
			r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			})

			adapter := chiadapter.New(r)

			req := events.APIGatewayProxyRequest{
				Path:       "/ping",
				HTTPMethod: "GET",
			}

			resp, err := adapter.Proxy(ctx, req)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
