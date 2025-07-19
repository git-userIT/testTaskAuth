package valid

import "regexp"

var regEm = regexp.MustCompile(`^[a-z][a-zA-Z0-9_]+@[a-z0-9]+\.[a-z]{2,}$`)

func CheckEmail(e string) bool {
	return regEm.Match([]byte(e))
}
