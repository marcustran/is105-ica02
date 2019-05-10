package main

import (
	//"fmt"
	"github.com/marcustran/is105-ica02/ascii"
)

func main() {

	//ascii.IterateOverASCIIStringLiteral("ascii.ascii")

	strLiteral := ascii.GetASCIIStringLiteral()

	ascii.GreetingASCII()
	ascii.IterateOverASCIIStringLiteral(strLiteral)

}
