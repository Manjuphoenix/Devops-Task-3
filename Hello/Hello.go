package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	fmt.Println(os.Args[1])
	maP := make(map[string]interface{})
	text := bufio.NewScanner(os.Stdin)
	for text.Scan(){
		inp := text.Text()
		maP[inp] = inp
		if inp == "-1" {
			break
		}
	}
	for k, n := range maP{
		fmt.Printf("%v\t%v\n", k, n)
	}
}