// gateway/main.go
package main

import (
	"log"

	"net/http"

	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux()

	if err := pb.RegisterShopServiceGraphql(mux); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(":8081", nil))
}
