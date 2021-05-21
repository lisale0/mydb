package token

var tokenStrings = map[string]string{
	"+": "PLUS",
	"=": handleEqual(),
}

func handleEqual() string {
	return "EQUAL"
}
