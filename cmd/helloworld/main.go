package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiadapter "github.com/serverless-plus/tencent-serverless-go/chi"
	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"
	"github.com/xqbumu/scf-demo/pkg/version"
)

const appName = "helloworld"

var chiFaas *chiadapter.ChiFaas

func init() {
	log.Println("Chi start")

	r := chi.NewRouter()
	r.Route("/"+appName, func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			buf, _ := json.Marshal(map[string]interface{}{
				"message": "Hello Serverless Chi",
				"query":   r.URL.Query().Get("q"),
			})
			w.Write(buf)
		})
		r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			v := version.GetBuildInfo()
			w.Write([]byte(v.ToString()))
		})
	})

	chiFaas = chiadapter.New(r)
}

// Handler serverless faas handler
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	return chiFaas.ProxyWithContext(ctx, req)
}

func main() {
	faas.Start(Handler)
}
