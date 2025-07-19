package valid

import (
	"log"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	strCh := []byte(`exaMple@example.ru`)
	res := regEm.Match(strCh)

	if !res {
		log.Fatal("Неверный формат почты: ", strCh)
	}
}
