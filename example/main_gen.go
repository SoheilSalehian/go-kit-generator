// *********************************************************************
// PARTIALLY GENERATED FILE -- DO NOT EDIT unless custom instrumentation
// Parameterized with:
//
// EndpointName = Test
// RequestName = TestRequest
// RequestType = int
// ResponseName = TestResponse
// ResponseType = string
// *********************************************************************

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

  // Setup logging
	logger := log.NewLogfmtLogger(os.Stderr)

  // Setup instrumentation
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounter(stdprometheus.CounterOpts{
  Namespace: "TestGroup",
		Subsystem: "[Add Microservice Name]",
		Name:      "requestCount",
		Help:      "Number of requests recieved",
	}, fieldKeys)

	requestLatency := metrics.NewTimeHistogram(time.Microsecond, kitprometheus.NewSummary(stdprometheus.SummaryOpts{
  Namespace: "TestGroup",
		Subsystem: "[Add Microservice Name]",
		Name:      "requestLatencyMicroseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys))

  // Add custom instrumentation interfaces
  //
  //

  // Wire everything together
  var svc TestService
	svc = service{}
	svc = loggingMiddleware{logger, svc}

  // Add the custom instrumentation to the instrumentationMiddleware
	svc = instrumentationMiddleware{requestCount, requestLatency, svc}

  handleTest := httptransport.NewServer(
		ctx,
		makeTestEndpoint(svc),
		decodeTestRequest,
		encodeResponse,
	)

  http.Handle("/test", handleTest)
  http.Handle("/metrics", stdprometheus.Handler())
  _ = logger.Log("msg", "HTTP", "addr=", os.Getenv("U_TEST_URL"))
  _ = logger.Log("err", http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("U_TEST_PORT")), nil), nil)
}

