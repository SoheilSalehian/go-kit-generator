// *******************************************************
// PARTIALLY GENERATED FILE -- Edit the indicated sections
// Parameterized with:
//
// EndpointName = GenerateOcrQuality
// RequestName = OcrText
// RequestType = string
// ResponseName = QualityMetric
// ResponseType = OcrQuality
// ********************************************************

package main

type GenerateOcrQualityService interface {
	CallGenerateOcrQuality(string) (OcrQuality, error)
}

type service struct{}

type OcrQuality struct {

	// Add in the json struct elements that the service will return back in its response.
	//
	//

}

func (service) CallGenerateOcrQuality(input string) (OcrQuality, error) {

	output := OcrQuality{}

	// Add in the main logic of the service
	//
	//

	return output, err
}
