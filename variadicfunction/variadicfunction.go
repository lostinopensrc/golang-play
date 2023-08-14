package variadicfunction

import "strings"

func Join_String(s ...string) string {
	return strings.Join(s, "@")
}
