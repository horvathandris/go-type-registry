/*
	Package parser
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package parser

import (
	"bufio"
	"github.com/horvathandris/go-type-registry/registry"
	"github.com/horvathandris/go-type-registry/util/sets"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type registryTemplateInput struct {
	Package        string
	TypeNamesSet   sets.StringSet
	TypeNamesSlice []string
}

var regTmplIn = registryTemplateInput{
	TypeNamesSet: sets.NewStringSet(),
}
var findPackageName = true

func StartFile(inFile string, outFile string) {
	parse(inFile)
	fillTemplate(outFile)
}

func StartDir(inDir string, outFile string) {
	for _, file := range *parseDir(inDir) {
		parse(file)
	}
	fillTemplate(outFile)
}

func parseDir(rootDir string) *[]string {
	var files []string
	err := filepath.Walk(rootDir, walkFunc(&files))
	if err != nil {
		log.Fatalln("could not read some files in directory")
	}
	return &files
}

func walkFunc(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() {
			log.Printf("Ignoring directory %v\n", info.Name())
			return nil
		} else if filepath.Ext(path) != ".go" {
			log.Printf("Ignoring non-go file %v\n", info.Name())
			return nil
		}
		*files = append(*files, path)
		return nil
	}
}

func parse(filename string) {
	log.Printf("Parsing file at %v", filename)
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
			regTmplIn.TypeNamesSet.Add(line)
			//regTmplIn.TypeNames = append(regTmplIn.TypeNames, line)
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

	regTmplIn.TypeNamesSlice = regTmplIn.TypeNamesSet.ToSlice()

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
