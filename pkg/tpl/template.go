package tpl

func DefaultTemplate() []byte {
	return []byte(`
{{ range .Data.Groups }}
Name = {{ .Name }}
{{ end }}
`)
}
