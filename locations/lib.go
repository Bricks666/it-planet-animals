package locations

import "strings"

func bodyContainsNull(body []byte) bool {
	return strings.Contains(string(body), "null")
}
