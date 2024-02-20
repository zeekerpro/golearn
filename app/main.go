package main

import (
	"fmt"

	"rsc.io/quote"

	"gok/packages/logger"
)

func main() {
	fmt.Println(quote.Go())

	logger.Log("use logger")

}
