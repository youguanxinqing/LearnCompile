package core

import (
	"bufio"
	"fmt"
	"os"
)

var (
	parser = NewParser()
	rtime  = NewRuntime()
)

func Loop() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">>")
		if code, err := scanner.ReadString('\n'); err != nil {
			fmt.Println(err)
		} else {
			interaction(code)
		}
	}
}

func interaction(code string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	root := parser.parse(code)
	if res := rtime.evaluate(root); res != nil {
		fmt.Printf("%v\n", *res)
	}
}
