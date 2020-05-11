package core

import "testing"

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
