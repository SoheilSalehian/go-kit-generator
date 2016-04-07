package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

type TransportMetadata struct {
	EndpointName string
	RequestName  string
	RequestType  string
	ResponseName string
	ResponseType string
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) Transport(writer io.Writer, metadata TransportMetadata) error {
	tmpl, err := g.template()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) template() (*template.Template, error) {

	resource, err := Asset("transport.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("template")
	return tmpl.Parse(string(resource))
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}
