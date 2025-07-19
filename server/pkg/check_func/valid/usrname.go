package valid

func CheckUsrname(un string) bool {
	return len([]rune(un)) >= 3
}
