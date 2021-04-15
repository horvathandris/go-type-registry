package go_type_registry

import (
	"bufio"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

var typeNames []string

func Parse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "type") && strings.HasSuffix(line, "struct {") {
			line = strings.TrimLeft(line, "type")
			line = strings.Split(line, "struct")[0]
			line = strings.TrimSpace(line)
			if unicode.IsUpper(rune(line[0])) {
				typeNames = append(typeNames, line)
			}
		}
	}

	out, err := os.Create("out.go")
	if err != nil {
		log.Fatal(err)
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	tmpl, _ := template.New("registry").Parse(registryTemplate)
	err = tmpl.Execute(out, typeNames)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}
