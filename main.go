package main

import (
	"capsulefm/libs/util/ugraphql"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("qlschema <url> [-debug]")
		return
	}
	inDataStruct := true
	debug := (len(os.Args) > 2 && os.Args[2] == "-debug")
	err := ugraphql.PrintSchemaFromUrl(os.Args[1], inDataStruct, debug)
	if err != nil {
		fmt.Println("err:", err)
	}
}
