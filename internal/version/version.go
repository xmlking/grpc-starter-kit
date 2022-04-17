package version

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed version.txt
	vcsTag      string
	vcsTime     string
	vcsRevision string
	//go:embed branch.txt
	vcsBranch   string
	vcsModified bool = true
	goVersion   string
	goCompiler  string
	goOS        string
	goArch      string
)

func init() {
	goVersion = runtime.Version()
	goCompiler = runtime.Compiler
	goOS = runtime.GOOS
	goArch = runtime.GOARCH

	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" {
			vcsTag = info.Main.Version
		}
	} else {
		log.Warn().Msg("not built in module mode")
	}
}

// BuildInfo show build status at *compile-time*
type BuildInfo struct {
	// GitVersion is populated from VCS. e.g., "v1.0.0" or "v1.1.2-dirty" or "v1.0.1-5-g585c78f-dirty" or "fbd157c"
	GitVersion string `json:"tag"`
	// GitCommit is populated from VCS. e.g., "86ec240af8cbd1b60bcc4c03c20da9b98005b92e"
	GitCommit string `json:"commit"`
	// GitBranch is populated from VCS
	GitBranch string `json:"branch"`
	// GitState will be "clean" or "dirty" based on if codebase is modified after commit
	GitState string `json:"state"`
	// GitBuildTime is populated from last VCS commit time. e.g., "2021-12-16T11:41:01Z"
	GitBuildTime string `json:"build_time"`
	// GoVersion is populated build-time go version. e.g.,  "go1.18"
	GoVersion string `json:"go_version"`
	// GoCompiler is populated build-time go compiler. e.g., "gc" or "gccgo"
	GoCompiler string `json:"compiler"`
	// GoPlatform is populated build-time go platform. e.g., "darwin/amd64"
	GoPlatform string `json:"platform"`
}

func (b BuildInfo) String() string {
	buf, err := json.Marshal(b)
	if err != nil {
		log.Fatal().Err(err).Msg("BuildInfo Marshal failed")
	}
	return string(buf)
}

func (b BuildInfo) MarshalZerologObject(e *zerolog.Event) {
	e.Str("tag", b.GitVersion).
		Str("commit", b.GitCommit).
		Str("branch", b.GitBranch).
		Str("state", b.GitState).
		Str("build_time", b.GitBuildTime).
		Str("go_version", b.GoVersion).
		Str("compiler", b.GoCompiler).
		Str("platform", b.GoPlatform)
}

// buildInfoTmpl is the message that is shown after process started.
const buildInfoTmpl = `
git tag         : %s
git build-time  : %s
git commit      : %s
git branch      : %s
git state       : %s
go version      : %s
go compiler     : %s
go platform     : %s
`

func (b BuildInfo) PrettyString() string {
	return fmt.Sprintf(buildInfoTmpl, b.GitVersion, b.GitBuildTime, b.GitCommit, b.GitBranch, b.GitState, b.GoVersion, b.GoCompiler, b.GoPlatform)
}

// GetBuildInfo helper
func GetBuildInfo() BuildInfo {
	toGitStatus := func(modified bool) string {
		if modified {
			return "dirty"
		} else {
			return "clean"
		}
	}
	return BuildInfo{
		strings.TrimSpace(vcsTag),
		vcsRevision,
		strings.TrimSpace(vcsBranch),
		toGitStatus(vcsModified),
		vcsTime,
		goVersion,
		goCompiler,
		fmt.Sprintf("%s/%s", goOS, goArch),
	}
}
