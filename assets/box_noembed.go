//go:build !embed

package assetsfs

import "embed"

var ebox embed.FS

func init() {
	IsEmbedded = false
}
