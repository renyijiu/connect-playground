package main

import (
	v1alpha1 "buf.build/gen/go/angry-popcorn/argo-workflows-apis/protocolbuffers/go/coscene/workflow/v1alpha1"
	"buf.build/gen/go/angry-popcorn/connect-playground/bufbuild/connect-go/proto/greet/v1/greetv1connect"
	v1 "buf.build/gen/go/angry-popcorn/connect-playground/protocolbuffers/go/proto/greet/v1"
	"context"
	"github.com/bufbuild/connect-go"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct {
}

func (ps *GreetServer) Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1alpha1.WorkflowTemplate], error) {
	return nil, nil
}

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(greetv1connect.NewGreetServiceHandler(&GreetServer{}))
	err := http.ListenAndServe(
		"localhost:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
