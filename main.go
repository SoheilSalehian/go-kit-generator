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
		fmt.Fprintln(os.Stderr, "  go-kit-generator -endpoint-name=<endpoint-name> -request-input=<request-input> -input-type=<input-type> -response-output=<response-output> -output-type=<output-type>")
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
	}

	flag.CommandLine.Init("", flag.ExitOnError)
}

func main() {
	endpointName := flag.String("endpoint-name", "", "Name of the endpoint based on the service")
	requestInput := flag.String("request-input", "", "")
	inputType := flag.String("input-type", "", "")
	responseOutput := flag.String("response-output", "", "")
	outputType := flag.String("output-type", "", "")

	flag.Parse()

	outputFile := "transport.go"
	writer, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	generator := &Generator{}

	m := metadata(*endpointName, *requestInput, *inputType, *responseOutput, *outputType)
	if err := generator.Generate(writer, m); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", outputFile)
}

func metadata(endpointName string, requestInput string, inputType string, responseOutput string, outputType string) (m Metadata) {
	m.EndpointName = endpointName
	m.RequestInput = requestInput
	m.InputType = inputType
	m.ResponseOutput = responseOutput
	m.OutputType = outputType

	return m
}
