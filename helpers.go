package main

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

var funcMap = template.FuncMap{
	"uppercase":  strings.ToUpper,
	"lowercase":  strings.ToLower,
	"nativetype": nativetype,
	"timestamp":  timestamp,
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func nativetype(input string) bool {
	nativeList := []string{"bool", "string", "int", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune", "float32", "float64", "complex64", "complex128"}
	for n := range nativeList {
		if input == nativeList[n] {
			// fmt.Println("Native type detected for request/response -- type: ", nativeList[n])
			return true
		}
	}
	return false
}

func timestamp() string {
	return time.Now().Format(time.RFC850)
}
