package main

import (
	"fmt"

	"github.com/singlaanish56/wcgo"
)
func main(){
	fmt.Println("Implementation of the unix command line tool wc")

	wcgo.InitFlags()
	wcgo.ParseTheFlags()
}