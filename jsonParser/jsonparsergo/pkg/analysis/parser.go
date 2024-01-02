package analysis

import (
	"bufio"
	"os"
	"fmt"
	"github.com/singlaanish56/jsonparsergo/pkg/errors"
)
var store []byte

func printTheTokens(input string){
	lexer := CreateLexer(input)
	for lexer.currentChar != 0{
		token:= lexer.GetToken()
		fmt.Printf("%s\n",token.Char)
	}

}
func ParseTheFile(fileName string) {

	f, err := os.Open(fileName)
	errors.HandleError(err)

	defer f.Close()

	reader := bufio.NewReader(f)
	
	for{
		line, _, err := reader.ReadLine();
		if errors.HandleFileError(err){
			break
		}

		store = append(store,line...)
	}
	
	printTheTokens(string(store))
}

func ParseTheString(stringName string) {

	printTheTokens(stringName)

}