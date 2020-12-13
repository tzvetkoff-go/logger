package logger

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONFormatter is a simple JSON log formatter.
type JSONFormatter struct {
}

// Format returns log fields formatted as a string.
func (f *JSONFormatter) Format(fields Fields) string {
	for k, v := range fields {
		if vv, ok := v.(error); ok {
			fields[k] = vv.Error() // Errors are ignored by encoding/json.
		}
	}

	bytes, err := json.Marshal(fields)
	if err != nil {
		return fmt.Sprintf(
			"{\"error\":\"%s\"}",
			strings.Replace(
				strings.Replace(
					err.Error(),
					"\\",
					"\\\\",
					-1,
				),
				"\"",
				"\\\"",
				-1,
			),
		)
	}

	return string(bytes)
}
