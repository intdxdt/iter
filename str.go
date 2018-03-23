package iter

import (
	"bytes"
	"fmt"
)

func IntArrayString(values []int) string {
	var buffer bytes.Buffer
	n := len(values) - 1

	buffer.WriteString("[")
	for i, o := range values {
		token := fmt.Sprintf("%v", o)
		if i < n {
			token += ", "
		}
		buffer.WriteString(token)
	}
	buffer.WriteString("]")
	return buffer.String()
}

