package tpl

func DefaultTemplate() []byte {
	return []byte(`
{{ range .Data.Groups }}
{{ .Name | upper }}
{{ end }}
`)
}
