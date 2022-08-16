package docs

import "embed"

//go:embed index.html
//go:embed openapi.yml
var DocsContent embed.FS
