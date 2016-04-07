package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
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

	tmpl := template.New("transportTemplate").Funcs(funcMap)
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

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

var funcMap = template.FuncMap{
	"uppercase": strings.ToUpper,
	"lowercase": strings.ToLower,
}
