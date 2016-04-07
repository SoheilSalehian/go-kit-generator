package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	EndpointName string `json:"endpoint_name"`
	RequestName  string `json:"request_name"`
	RequestType  string `json:"request_type"`
	ResponseName string `json:"response_name"`
	ResponseType string `json:"response_type"`
}

func main() {
	raw, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	var conf Configuration
	json.Unmarshal(raw, &conf)

	configuration := fmt.Sprintf("package main\n//go:generate ../go-kit-generator -endpoint-name=%s -request-name=%s -request-type=%s -response-name=%s -response-type=%s \n", conf.EndpointName, conf.RequestName, conf.RequestType, conf.ResponseName, conf.ResponseType)

	err = ioutil.WriteFile(fmt.Sprintf("%s/gen.go", os.Args[2]), []byte(configuration), 0644)
	if err != nil {
		panic(err)
	}
}
