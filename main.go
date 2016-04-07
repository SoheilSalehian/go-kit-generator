package main

import (
	"flag"
	"fmt"
	"os"
)

var endpointName = flag.String("endpoint-name", "", "Name of the endpoint based on the service")
var requestName = flag.String("request-name", "", "")
var requestType = flag.String("request-type", "", "")
var responseName = flag.String("response-name", "", "")
var responseType = flag.String("response-type", "", "")

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of go-kit-generator:")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "  go-kit-generator -endpoint-name=<endpoint-name> -request-name=<request-name> -request-type=<request-type> -response-name=<response-name> -response-type=<response-type>")
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
	}

	flag.CommandLine.Init("", flag.ExitOnError)

	flag.Parse()
}

func transportGen() {

	outputFile := "transport.go"
	writer, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	generator := &Generator{}

	m := Metadata{*endpointName, *requestName, *requestType, *responseName, *responseType}
	if err := generator.Transport(writer, m); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", outputFile)
}

func serviceGen() {

	outputFile := "service.go"
	writer, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	generator := &Generator{}

	m := Metadata{*endpointName, *requestName, *requestType, *responseName, *responseType}
	if err := generator.Service(writer, m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Printf("Generated %s\n", outputFile)
}

func main() {
	transportGen()
	serviceGen()
}
