package tpl

func DefaultTemplate() []byte {
	return []byte(`
{{- range .Data.Groups -}}
# {{ .Name | upper }}
{{- range .Rules }}
## {{ .Name }}
### Query:
{{ .Query }}
### Duration: {{ div .Duration 60 }}m
### Labels:
{{- range $key, $value :=  .Labels }}
>  {{ $key }}: {{ $value }}
{{- end }}
### Annotations:
{{- range $key, $value :=  .Annotations }}
>  {{ $key }}: {{ $value }}
{{- end }}
{{- end }}
{{- end }}
`)
}
