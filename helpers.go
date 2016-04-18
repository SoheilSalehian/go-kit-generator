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
	"zero":       zero,
	"s3type":     s3type,
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func nativetype(input string) bool {
	nativeList := []string{"bool", "string", "int", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune", "float32", "float64", "complex64", "complex128"}
	for n := range nativeList {
		if input == nativeList[n] {
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

func length(list []string) int {
	return len(list)
}

func last(list []string, i int) bool {
	fmt.Println(len(list), i)
	return len(list) == i
}

func zero(index int) bool {
	if index == 0 {
		return true
	}
	return false
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("service-order flag already set")
	}
	for _, svc := range strings.Split(value, ",") {
		*i = append(*i, svc)
	}
	return nil
}

func s3type(t string) bool {
	return t == "aws.S3Type"

}
