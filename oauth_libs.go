package eliotlibs

import "strings"

func IsAdminPath(path string) bool {
	return strings.HasPrefix(path, "/admin/")
}
