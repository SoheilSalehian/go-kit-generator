// *******************************************************
// PARTIALLY GENERATED FILE -- Edit the indicated sections
// Parameterized with:
//
// EndpointName = GenerateOcrQuality
// RequestName = OcrText
// RequestType = string
// ResponseName = QualityMetric
// ResponseType = OcrQuality
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

	handleGenerateOcrQuality := httptransport.NewServer(
		ctx,
		makeGenerateOcrQualityEndpoint(svc),
		decodeGenerateOcrQualityRequest,
		encodeResponse,
	)

	http.Handle("/generateocrquality", handleGenerateOcrQuality)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("GENERATEOCRQUALITY_PORT"))))
}
