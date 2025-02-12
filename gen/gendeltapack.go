package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

func main() {
	log.Println("gendeltapack start")

	ftmpl, err := os.ReadFile("gen/deltapack.tmpl")
	if err != nil {
		log.Println("failed to read template", err)
		os.Exit(1)
	}

	tmpl, err := template.New("deltapack").Parse(string(ftmpl))
	if err != nil {
		log.Println("failed to process template", err)
		os.Exit(2)
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, &deltapack{})
	if err != nil {
		log.Println("failed to execute template", err)
		os.Exit(3)
	}

	if err := os.WriteFile("deltapack_gen.go", out.Bytes(), 0600); err != nil {
		log.Println("failed to write file", err)
		os.Exit(4)
	}

	log.Println("gen-dbp end")
}
