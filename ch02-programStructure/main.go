package main

import (
	"fmt"

	"github.com/tinmanjk/tgpl/ch01-tutorial/02-commandLineArguments/exportExample"
)

func main() {
	sum := exportExample.Sum(3, 3)
	fmt.Println(sum)
}