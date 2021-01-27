package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	err := executeTemplate(os.Stdin, os.Stdout, os.Args[1:]...)
	if err != nil {
		panic(err)
	}
}

func executeTemplate(valuesIn io.Reader, output io.Writer, tplFiles ...string) error {
	tpl, err := template.ParseFiles(tplFiles...)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, valuesIn)
	if err != nil {
		return fmt.Errorf("failed to read standard input: %v", err)
	}

	values := make(map[string]interface{})
	err = yaml.Unmarshal(buf.Bytes(), &values)
	if err != nil {
		return fmt.Errorf("failed to parse yaml: %v", err)
	}

	err = tpl.Execute(output, values)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}
