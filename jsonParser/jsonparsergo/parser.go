package jsonparsergo

import (
	"bufio"
	"os"
	"fmt"
)
var store []byte

func parseObject
func checkIfJsonValid() int{
	fmt.Println(string(store))
	for i, _ := range store{
		if store[i]=='{'{
			i=parseObject(i,)
		}
	}
}

func ParseTheFile(fileName string) {

	f, err := os.Open(fileName)
	HandleError(err)

	defer f.Close()

	reader := bufio.NewReader(f)
	
	for{
		line, _, err := reader.ReadLine();
		if HandleFileError(err){
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