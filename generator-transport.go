package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

type Metadata struct {
	EndpointName   string
	RequestInput   string
	InputType      string
	ResponseOutput string
	OutputType     string
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
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
