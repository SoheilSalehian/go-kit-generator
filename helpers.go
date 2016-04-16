package main

import (
	"errors"
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
	"first":      first,
	"last":       last,
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

func first(list []string) string {
	return list[0]
}

func last(list []string) string {
	return list[len(list)-1]
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

// func (i *arrayFlags) Set(value string) error {
// 	*i = append(*i, value)
// 	return nil
// }

func (i *arrayFlags) Set(value string) error {
	// If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	if len(*i) > 0 {
		return errors.New("service-order flag already set")
	}
	for _, svc := range strings.Split(value, ",") {
		*i = append(*i, svc)
	}
	return nil
}
