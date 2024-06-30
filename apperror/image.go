package apperror

import (
	"fmt"
	"strings"
)

func NewImageSizeExceeded(maxSize string) error {
	return NewAppError(
		CodeBadRequest,
		fmt.Sprintf("image file size must be less than %s", maxSize),
		nil,
	)
}

func NewRestrictredFileType(types ...string) error {
	sb := strings.Builder{}
	sb.WriteString("file type must be ")
	for i := 0; i < len(types); i++ {
		if i == len(types)-1 {
			sb.WriteString(fmt.Sprintf("or %s", types[i]))
			continue
		}
		sb.WriteString(fmt.Sprintf("%s, ", types[i]))
	}

	return NewAppError(
		CodeBadRequest,
		sb.String(),
		nil,
	)
}
