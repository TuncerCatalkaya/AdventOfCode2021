package filereader

import (
	"os"
	"strings"
)

func ReadFile(path, split string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), split)
}
