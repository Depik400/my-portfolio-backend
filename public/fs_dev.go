//go:build !prod

package public

import (
	"io/fs"
	"os"
)

const folder = "./public/frontend/dist"

func GetFS() (fs.FS, error) {
	return os.DirFS(folder), nil
}
