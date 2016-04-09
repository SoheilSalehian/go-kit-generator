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
CallTest (int) (string, error)
}

type service struct{}

// Define the service output struct if not native type
 



func (service) CallTest(input int) (string, error) {

 
  output := ""


// Add in the main logic of the service
//
//

	return output, nil
}
