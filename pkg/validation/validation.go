package validation

import (
	"errors"
	"strings"
	"unicode/utf8"
)

func ValidSegmentName(segment string) (bool, error) {
	count := utf8.RuneCountInString(segment)

	if count < 6 {
		return false, errors.New("the minimum segment length is 6 characters")
	}

	if !strings.HasPrefix(segment, "AVITO") {
		return false, errors.New("the first word must be AVITO")

	}

	if segment[5] != '_' {
		return false, errors.New(`after word AVITO must be symbol "_"`)
	}

	return true, nil
}
