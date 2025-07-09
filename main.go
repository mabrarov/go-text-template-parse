package main

import (
	"fmt"
	"html/template"
	"text/template/parse"
)

func main() {
	t := template.Must(template.New("cooltemplate").Parse(`<h1>{{ .name }} {{ printf "%d" .age }}</h1>`))
	fmt.Println(listTemplateFields(t))
}

func listTemplateFields(t *template.Template) []string {
	return listNodeFields(t.Tree.Root)
}

func listNodeFields(node parse.Node) []string {
	var res []string
	if node.Type() == parse.NodeAction {
		res = append(res, node.String())
	}

	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			res = append(res, listNodeFields(n)...)
		}
	}
	return res
}
