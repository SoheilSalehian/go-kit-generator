package main

import (
	"fmt"
	"html/template"
	"strings"
)

var funcMap = template.FuncMap{
	"uppercase": strings.ToUpper,
	"lowercase": strings.ToLower,
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}
