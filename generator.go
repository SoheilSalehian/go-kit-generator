package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

type Metadata struct {
	EndpointName string
	RequestName  string
	RequestType  string
	ResponseName string
	ResponseType string
}

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

	tmpl := template.New("transportTemplate")
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

	tmpl := template.New(templateFile)
	return tmpl.Parse(string(resource))
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}
