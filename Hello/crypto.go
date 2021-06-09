package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	amount:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter the amount to be added")
	println(amount)

}
