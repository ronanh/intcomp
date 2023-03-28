package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

func main() {
	var out bytes.Buffer
	log.Println("gendeltapack start")

	data := []deltapackint{
		{32, false, false},
		{32, true, false},
		{64, false, true},
		{64, true, false},
	}

	ftmpl, err := os.ReadFile("gen/deltapackint.tmpl")
	if err != nil {
		log.Println("failed to read template", err)
		os.Exit(1)
	}

	tmpl, err := template.New("deltapackint").Parse(string(ftmpl))
	if err != nil {
		log.Println("failed to process template", err)
		os.Exit(2)
	}

	err = tmpl.Execute(&out, data)
	if err != nil {
		log.Println("failed to execute template", err)
		os.Exit(3)
	}

	if err := os.WriteFile("deltapackint.go", out.Bytes(), 0600); err != nil {
		log.Println("failed to write file", err)
		os.Exit(4)
	}

	log.Println("gen-dbp end")
}
