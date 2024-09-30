package main

type CRD struct {
	claimName string
	Fields []Field
}

type Field struct {
	Type string
	Title string
	Description string
	Required bool
}
