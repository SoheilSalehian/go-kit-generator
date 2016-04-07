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

type TestService interface {
	CallTest(int) (string, error)
}

type service struct{}

type string struct {

	// Add in the json struct elements that the service will return back in its response.
	//
	//

}

func (service) CallTest(input int) (string, error) {

	output := string{}

	// Add in the main logic of the service
	//
	//

	return output, err
}
