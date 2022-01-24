package version

import (
	"bytes"
	"runtime/debug"
	"text/template"

	"github.com/rs/zerolog/log"
)

var bomTmpl = ` mod	{{.Main.Path}}	{{.Main.Version}}	{{.Main.Sum}}
{{range .Deps}} dep	{{.Path}}	{{.Version}}	{{.Sum}}{{if .Replace}}
	=> {{.Replace.Path}}	{{.Replace.Version}}	{{.Replace.Sum}}{{end}}
{{end}}`

func GetSoftwareBOM() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Warn().Msg("not built in module mode")
		return "not built in module mode"
	}

	buf := new(bytes.Buffer)
	err := template.Must(template.New("bom").Parse(bomTmpl)).Execute(buf, info)
	if err != nil {
		log.Fatal().Err(err).Msg("GetSoftwareBOM bomTmpl mapping failed")
	}
	return buf.String()
}
