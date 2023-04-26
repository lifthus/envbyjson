package main

import (
	"fmt"
	"os"

	envjson "github.com/lifthus/envjson/go"
)

func main() {

	err := envjson.LoadProp("c")
	if err != nil {
		panic(err)
	}

	fmt.Printf("a=%s\n", os.Getenv("a"))
	fmt.Printf("b=%s\n", os.Getenv("b"))
	fmt.Printf("d=%s\n", os.Getenv("d"))
	fmt.Printf("e=%s\n", os.Getenv("e"))
	fmt.Printf("f=%s\n", os.Getenv("f"))
	fmt.Printf("g=%s\n", os.Getenv("g"))

}
