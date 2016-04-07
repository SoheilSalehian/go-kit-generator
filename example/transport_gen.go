// *********************************************
// GENERATED FILE -- DO NOT EDIT PLEASE
// Parameterized with:
//
// EndpointName = Test
// RequestName = TestRequest
// RequestType = int
// ResponseName = TestResponse
// ResponseType = string
// *********************************************

package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// Request
type TestRequest struct {
	TestRequest int `json:"TestRequest"`
}

// Response
type TestResponse struct {
	TestResponse string `json:"TestResponse"`
	Err          string `json:"err,omitempty"`
}

// Endpoint(s)
func makeTestEndpoint(svc TestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TestRequest)
		serviceOutput, err := svc.Test(req.TestRequest)
		if err != nil {
			return TestResponse{serviceOutput, err}, nil
		}
		return TestResponse{serviceOutput, ""}, nil
	}
}

// Decode Request
func decodeTestRequest(r *http.Request) (interface{}, error) {
	var request TestRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Encode Response
func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
