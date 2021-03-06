package main

import (
	"errors"
	"io"
	"text/template"
)

const (
	JSON = "json"
)

type Metadata struct {
	PackageName string
	Type        string
}

type Generator struct {
	Format string
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.template()
	if err != nil {
		return err
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) template() (*template.Template, error) {
	if g.Format != JSON {
		return nil, errors.New("Unsupported format")
	}

	//tmpl, e := template.ParseFiles("/Users/vikas.verma/go/src/github.com/vikasverma155/go-fun/jsongen/tmpl/write_to_json.tmpl")
	tmpl := template.New("jsonTemplate")
	tmpl, e := tmpl.Parse(templateString())
	return tmpl, e
}

func templateString() string {
	return `//This is Aman's Generated File
//Request you not to mess with it :)

package {{ .PackageName }}

import (
	"encoding/json"
	"io"
)

func (obj {{ .Type }}) WriteTo(writer io.Writer) (int64, error) {
	data, err := json.Marshal(&obj)
	if err != nil {
		return 0, err
	}
	length, err := writer.Write(data)
	return int64(length), err
}`
}
