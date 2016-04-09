package main

import (
	"bytes"
	"html/template"
	"io"
)

type Metadata struct {
	EndpointName string
	RequestName  string
	RequestType  string
	ResponseName string
	ResponseType string
}

var NativeType bool

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) Transport(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.transporTemplate()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) transporTemplate() (*template.Template, error) {

	templateFile := "templates/transport.tmpl"

	resource, err := Asset(templateFile)
	if err != nil {
		return nil, err
	}

	tmpl := template.New(templateFile).Funcs(funcMap)
	return tmpl.Parse(string(resource))
}

func (g *Generator) Service(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.serviceTemplate()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) serviceTemplate() (*template.Template, error) {

	templateFile := "templates/service.tmpl"

	resource, err := Asset(templateFile)
	if err != nil {
		return nil, err
	}

	tmpl := template.New(templateFile).Funcs(funcMap)
	return tmpl.Parse(string(resource))
}

func (g *Generator) Main(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.mainTemplate()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) mainTemplate() (*template.Template, error) {

	// NOTE: main.tmpl makes go-bindata not parse the template properly
	templateFile := "templates/srv-main.tmpl"

	resource, err := Asset(templateFile)
	if err != nil {
		return nil, err
	}

	tmpl := template.New(templateFile).Funcs(funcMap)
	return tmpl.Parse(string(resource))
}

func (g *Generator) loggingTemplate() (*template.Template, error) {

	templateFile := "templates/logging.tmpl"

	resource, err := Asset(templateFile)
	if err != nil {
		return nil, err
	}

	tmpl := template.New(templateFile).Funcs(funcMap)
	return tmpl.Parse(string(resource))
}

func (g *Generator) Logging(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.loggingTemplate()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) instrumentationTemplate() (*template.Template, error) {

	templateFile := "templates/instrumentation.tmpl"

	resource, err := Asset(templateFile)
	if err != nil {
		return nil, err
	}

	tmpl := template.New(templateFile).Funcs(funcMap)
	return tmpl.Parse(string(resource))
}

func (g *Generator) Instrumentation(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.instrumentationTemplate()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}
