// *******************************************************
// PARTIALLY GENERATED FILE -- Edit the indicated sections
// Parameterized with:
//
// EndpointName = Test
// RequestName = TestRequest
// RequestType = int
// ResponseName = TestResponse
// ResponseType = string
// ********************************************************

package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	svc := service{}

	handleTest := httptransport.NewServer(
		ctx,
		makeTestEndpoint(svc),
		decodeTestRequest,
		encodeResponse,
	)

	http.Handle("/test", handleTest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("TEST_PORT"))))
}
