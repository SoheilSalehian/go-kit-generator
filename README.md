# go-kit-generator
Generate the basic parts of a micro-service written using go-kit

# Usage
As each microservice requires roughly the same boilerplate code, using the go-kit-generator executable, the initial code and Dockerfile can be generated to efficiently get you going in writing the microservice. The user needs to specify the following to the go-kit-generator tool:

$ ./go-kit-generator -help
Usage of go-kit-generator:

  go-kit-generator -endpoint-name=<endpoint-name> -request-name=<request-name> -request-type=<request-type> -response-name=<response-name> -response-type=<response-type>

  -endpoint-name string
        Name of the endpoint for the microservice
  -request-name string
        Name of the request to the microservice
  -request-type string
        Request type of the microservice
  -response-name string
        Name of the response to the microservice
  -response-type string
        Response type of the microservice
