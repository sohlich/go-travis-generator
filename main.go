package main

import (
	"log"
	"os/exec"
	"strings"

	"text/template"

	"os"
)

var (
	imports  = `{{range .Imports }} {{.}}{{end}}`
	standart = `{{.Standard}}`

	travisTempl = `
language: go
go:
	- tip
		
notifications:
	email: false

install:
	{{range .Results }}
	-{{.}}{{end}}
`
)

func main() {
	listCmd := exec.Command("go", "list", "-f", imports)
	o, err := listCmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}

	all := strings.Fields(string(o))

	sLst := strings.Fields(string(o))
	standardMap := make(map[string]bool)
	for _, val := range sLst {
		standardMap[val] = true
	}

	result := make([]string, 0)
	for _, val := range all {
		listCmd = exec.Command("go", "list", "-f", standart, val)
		o, err = listCmd.CombinedOutput()
		if strings.TrimSpace(string(o)) == "false" {
			result = append(result, val)
		}
	}

	tmpl := template.New("out")
	tmpl, _ = tmpl.Parse(travisTempl)

	tmpl.Execute(os.Stdout, struct {
		Results []string
	}{
		result,
	})
}
