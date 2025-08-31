//go:build embed

package assetsfs

import "embed"

//go:embed dist/*
var ebox embed.FS

func init() {
	IsEmbedded = true
}
