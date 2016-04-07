// GENERATED FILE -- DO NOT EDIT PLEASE
package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// Request
type GenerateOcrQualityRequest struct {
OcrText string `json:"OcrText"`
}

// Response
type GenerateOcrQualityResponse struct {
  QualityMetric OcrQuality `json:"QualityMetric"`
	Err    string            `json:"err,omitempty"`
}

// Endpoint(s)
func makeGenerateOcrQualityEndpoint(svc GenerateOcrQualityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
  req := request.(GenerateOcrQualityRequest)
  output, err := svc.GenerateOcrQuality(req.OcrText)
		if err != nil {
    return GenerateOcrQualityResponse{\QualityMetric, err.Error()}, nil
		}
    return GenerateOcrQuality Response{\QualityMetric, ""}, nil
	}
}

// Decode Request
func decodeGenerateOcrQualityRequest(r *http.Request) (interface{}, error) {
var request GenerateOcrQualityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Encode Response
func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

