package template

const ModelTemplate = `
type {{.Model}} struct {
	{{- range $i, $v := .Fields}}
	{{$v.Key}} {{$v.FieldType}} {{$v.FieldName}}
	{{- end}}
}

func (p *{{.Model}}) TableName() string {
	return "{{.Table}}"
}
`
