package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
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

	for _, v := range data {
		var out bytes.Buffer
		err = tmpl.Execute(&out, &v)
		if err != nil {
			log.Println("failed to execute template", err)
			os.Exit(3)
		}
		filename := fmt.Sprintf("deltapack%s.go", strings.ToLower(v.Typename()))

		if err := os.WriteFile(filename, out.Bytes(), 0600); err != nil {
			log.Println("failed to write file", err)
			os.Exit(4)
		}
	}

	log.Println("gen-dbp end")
}
