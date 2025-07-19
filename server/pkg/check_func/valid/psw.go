package valid

func CheckPassLen(p string) bool {
	return len([]rune(p)) >= 6
}
