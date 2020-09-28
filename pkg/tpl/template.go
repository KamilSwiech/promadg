package tpl

func DefaultTemplate() []byte {
	return []byte(`
{{ range .Data.Groups }}
{{ .Name | upper }}
{{ range .Rules }}
{{ .Name }}
Query: {{ .Query }}
Duration: {{ .Duration }}
{{ end }}
{{ end }}
`)
}
