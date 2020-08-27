package tpl

func template () []byte {
	return []byte(`
{{ range .Data.Groups }}
Name = {{ .Name }}
{{ end }}
`)
