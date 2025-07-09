package main

import (
	"fmt"
	"html/template"
	"slices"
	"text/template/parse"
)

func main() {
	t := template.Must(template.New("cooltemplate").Parse(`<h1>{{ .name }} {{ printf "%d" .age }}</h1>`))
	fmt.Println(listTemplateFields(t.Tree.Root))
}

func listTemplateFields(node parse.Node) []string {
	var ids []string
	//goland:noinspection ALL
	switch node.Type() {
	case parse.NodeList:
		listNode := node.(*parse.ListNode)
		for _, n := range listNode.Nodes {
			ids = append(ids, listTemplateFields(n)...)
		}
	case parse.NodeAction:
		actionNode := node.(*parse.ActionNode)
		if actionNode.Pipe == nil {
			break
		}
		for _, cmdNode := range actionNode.Pipe.Cmds {
			ids = append(ids, listTemplateFields(cmdNode)...)
		}
	case parse.NodeCommand:
		commandNode := node.(*parse.CommandNode)
		for _, node := range commandNode.Args {
			ids = append(ids, listTemplateFields(node)...)
		}
	case parse.NodeField:
		fieldNode := node.(*parse.FieldNode)
		ids = slices.Clone(fieldNode.Ident)
	}
	return ids
}
