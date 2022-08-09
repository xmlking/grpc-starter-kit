package embed

import "embed"

//go:embed config/certs/*.pem config/config*.yml
var StaticConfig embed.FS
