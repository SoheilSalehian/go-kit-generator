// ****************************************************************
// GENERATED FILE -- DO NOT EDIT Unless Response type is not native
// Parameterized with:
//
// EndpointName = Test
// RequestName = TestRequest
// RequestType = int
// ResponseName = TestResponse
// ResponseType = string
// ****************************************************************

package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
  TestService
}


// Logging decoupled from the service logic
func (mw loggingMiddleware) Test(TestRequest int) (output string, err error) {
  defer func(begin time.Time) {
    _ = mw.logger.Log(
    "method", "Test",
    "input", TestRequest,
    // NOTE: Ensure type "string" is printable here
    "output", output,
    "err", err,
    "took", time.Since(begin),
    )
  }(time.Now())

  output, err = mw.TestService.CallTest(TestRequest)
  if err != nil {
   
    return "", err
  
  }

  return output, nil
}



