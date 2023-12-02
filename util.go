package advent

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(file string) string {
	dat, err := os.ReadFile(file)
	check(err)
	return string(dat)
}
