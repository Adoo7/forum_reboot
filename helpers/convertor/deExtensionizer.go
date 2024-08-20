package convertor

import (
	"errors"
	"strings"
)

func TrimExtension(s string) (string, error) {
	fileNameAndSuffix := strings.Split(s, ".")
	if len(fileNameAndSuffix) != 2 {
		println("error trimming file")
		return s, errors.New("error trimming extension, file should be in following format: 'file.extension'")
	}
	return fileNameAndSuffix[0], nil
}
