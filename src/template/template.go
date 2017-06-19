package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Friends []*Friend
}

type Friend struct {
	Name string
	Age  uint8
}

func main() {
	Li := Friend{"Li", 18}
	Wang := Friend{"Wang", 18}
	t := template.New("example")
	t, _ = t.Parse(`
	Hello , {{.Name}}!
	{{range .Friends}}
	My friend name is {{.Name}}, {{if .Age}} years old this year {{end}}.
	{{end}}`)
	me := Person{Name: "SheJie", Friends: []*Friend{&Li, &Wang}}
	t.Execute(os.Stdout, me)
}
