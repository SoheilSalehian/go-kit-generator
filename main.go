package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of go-kit-generator:")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "  go-kit-generator -endpoint-name=<endpoint-name> -request-name=<request-name> -request-type=<request-type> -response-name=<response-name> -response-type=<response-type>")
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
	}

	flag.CommandLine.Init("", flag.ExitOnError)
}

func transportGen() {
	endpointName := flag.String("endpoint-name", "", "Name of the endpoint based on the service")
	requestName := flag.String("request-name", "", "")
	requestType := flag.String("request-type", "", "")
	responseName := flag.String("response-name", "", "")
	responseType := flag.String("response-type", "", "")

	flag.Parse()

	outputFile := "transport.go"
	writer, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	generator := &Generator{}

	m := transportMetadata(*endpointName, *requestName, *requestType, *responseName, *responseType)
	if err := generator.Transport(writer, m); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", outputFile)
}

func transportMetadata(endpointName string, requestName string, requestType string, responseName string, responseType string) (m TransportMetadata) {
	m.EndpointName = endpointName
	m.RequestName = requestName
	m.RequestType = requestType
	m.ResponseName = responseName
	m.ResponseType = responseType

	return m
}

func main() {
	transportGen()
}
