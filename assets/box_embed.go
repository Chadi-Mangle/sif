//go:build embed

package assetsfs

import "embed"

//go:embed dist/*.css
var ebox embed.FS

func init() {
	IsEmbedded = true
}
