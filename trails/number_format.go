package trails

import (
	"fmt"
	"strings"
)

func NumberFormat(number int) string {
	strNumber := fmt.Sprintf("%d", number)
	parts := make([]string, 0)

	for len(strNumber) > 3 {
		parts = append(parts, strNumber[len(strNumber)-3:])
		strNumber = strNumber[:len(strNumber)-3]
	}
	parts = append(parts, strNumber)

	// Reverse the order of parts
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	formattedNumber := strings.Join(parts, ",")
	return formattedNumber
}
