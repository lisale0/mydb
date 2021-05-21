package parser

type Interpreter struct {
}

type Token struct {
}

type Scanner struct {
	interpreter Interpreter
	source      string
	tokens      []Token
}

func NewScanner(interpreter Interpreter, source string) *Scanner {
	return &Scanner{
		interpreter: interpreter,
		source:      source,
		tokens:      []Token{},
	}
}

func (s *Scanner) ScanToken() {

}

var tokenStrings = map[string]string{
	"+": "PLUS",
	"=": handleEqual(),
}

func handleEqual() string {
	return "EQUAL"
}
