package core

import (
	"fmt"
	"testing"
)

func TestInteraction_loop(t *testing.T) {
	Loop()
}

func TestScan(t *testing.T) {
	var code string
	for {
		//fmt.Printf(">>")
		if _, err := fmt.Scanln(&code); err != nil {
			fmt.Printf(code)
		}
	}
}