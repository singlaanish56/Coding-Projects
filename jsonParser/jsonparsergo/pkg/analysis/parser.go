package analysis

import (
	"bufio"
	"os"
	"fmt"
	"github.com/singlaanish56/jsonparsergo/pkg/errors"
)
var store []byte

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

	fmt.Printf("%d\t", checkIfJsonValid())
}

func ParseTheString(stringName string) {

	store = []byte(stringName)

	fmt.Printf("%d\t", checkIfJsonValid())
}