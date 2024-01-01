package tokens

const (
	String = "STR"
	Number = "NUM"
	True   = "T"
	False  = "F"
	Null   = "N"

	LeftCurlyBR  = "{"
	RightCurlyBR = "}"
	LeftBoxBR    = "["
	RightBoxBR   = "]"

	Colon = ":"
	Comma = ","
)

type Token struct {
	Char     string
	Name     string
	TokenStart int
	TokenEnd int
}
