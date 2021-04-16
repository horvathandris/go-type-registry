/*
	Package parser
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package parser

import (
	"bufio"
	"github.com/horvathandris/go-type-registry/registry"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

type registryTemplateInput struct {
	Package   string
	TypeNames []string
}

var regTmplIn registryTemplateInput
var findPackageName = true

func Start(inFile string, outFile string) {
	parse(inFile)
	fillTemplate(outFile)
}

func parse(filename string) {
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
		if findPackageName && strings.HasPrefix(scanner.Text(), "package") {
			regTmplIn.Package = strings.TrimSpace(strings.TrimLeft(scanner.Text(), "package"))
			findPackageName = false
		} else if line, ok := checkIfStructInitLine(scanner.Text()); ok {
			regTmplIn.TypeNames = append(regTmplIn.TypeNames, line)
		}
	}
}

func fillTemplate(filename string) {
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	tmpl := template.Must(template.New("registry").Funcs(funcMap).Parse(registry.Template))
	err = tmpl.Execute(out, regTmplIn)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}

func checkIfStructInitLine(line string) (string, bool) {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "type") && strings.HasSuffix(line, "struct {") {
		line = strings.TrimSpace(strings.Split(strings.TrimLeft(line, "type"), "struct")[0])
		if unicode.IsUpper(rune(line[0])) {
			return line, true
		}
	}
	return "", false
}
