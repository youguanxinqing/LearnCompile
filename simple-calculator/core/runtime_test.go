package core

import (
	"fmt"
	"testing"
)

func TestRuntime_evaluate(t *testing.T) {
	parser := NewParser()
	//root := parser.parse("int a=1 + 1; a = a + 10; a")
	root := parser.parse(` 
		int a; 
		int b = 2 / 2; 
		a = 10; 
		a + b / 10;
		10;
		a = a + 12 * 3 - 1; 
		a + b;
		b = a;
	`)
	runtime := NewRuntime()
	res := runtime.evaluate(root)
	if res != nil {
		t.Log(*res)
	} else {
		t.Log(res)
	}
}

func TestRuntime_evaluate_parenthese(t *testing.T) {
	parser := NewParser()
	//root := parser.parse(`2 * (3 + 5) / 3`)
	root := parser.parse(`(2 * 5) / 3`)
	runtime := NewRuntime()
	ret := runtime.evaluate(root)
	if ret != nil {
		t.Log(*ret)
	}
}

func Test_a(t *testing.T) {
	for i := range []int32{1, 2, 3, 4} {
		fmt.Println(i)
	}
}
