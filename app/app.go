package app

import (
	"context"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sodefrin/appengine-boilerplate/app/controller"
	"github.com/sodefrin/appengine-boilerplate/proto"
)

// Run app.
func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcHandler := runtime.NewServeMux()
	err := proto.RegisterAppengineBoilerplateHandlerServer(ctx, grpcHandler, controller.NewAppengineBoilerplate())
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", grpcHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if hasSuffix(r.URL.Path, ".js/.css/.ico/.png/.jpeg") {
			http.ServeFile(w, r, "web/dist"+r.URL.Path)
			return
		}
		http.ServeFile(w, r, "web/dist")
	})

	return http.ListenAndServe(":8080", mux)
}

func hasSuffix(path, suffix string) bool {
	ss := strings.Split(suffix, "/")
	for _, v := range ss {
		if ok := strings.HasSuffix(path, v); ok {
			return true
		}
	}
	return false
}
