//go:build go1.19

package version

import (
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Warn().Msg("not built in module mode")
		return
	}
	goVersion = info.GoVersion
	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			vcsRevision = kv.Value
		case "vcs.time":
			vcsTime = kv.Value
		case "vcs.modified":
			vcsModified = kv.Value == "true"
		case "-compiler":
			goCompiler = kv.Value
		case "GOOS":
			goOS = kv.Value
		case "GOARCH":
			goArch = kv.Value
		}
	}
}
