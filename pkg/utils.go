package pkg

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

//FatalF is a convinience function to print erros and exit
func FatalF(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(2)
}

//CreateFile is a function to create a file from a given template
func CreateFile(filePath string, fileTemplate string, values interface{}) {
	file, err := os.Create(filePath)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	tmpl, err := template.New("test").Parse(fileTemplate)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	err = tmpl.Execute(file, values)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func GetPekkaTraefikPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	return filepath.Join(cwd, "pekka-traefik")
}
