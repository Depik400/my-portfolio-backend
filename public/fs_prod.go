//go:build prod

package public

import (
	"embed"
	"io/fs"
)

//go:embed frontend/dist/*
var content embed.FS

func GetFS() (fs.FS, error) {
	stripped, err := fs.Sub(content, "frontend/dist")
	return stripped, err
}
