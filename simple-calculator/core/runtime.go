package core

import (
	"LearnCompile/simple-calculator/core/ast"
	"fmt"
	"strconv"
)

type Store struct {
	data1 map[string]int      // 声明, 初始化
	data2 map[string]struct{} // 声明,未初始化
}

func NewStore() *Store {
	return &Store{
		data1: make(map[string]int),
		data2: make(map[string]struct{}),
	}
}

// 声明变量
func (s *Store) declare(key string) error {
	if _, ok := s.data2[key]; ok {
		return fmt.Errorf("变量 '%s' 重复声明", key)
	}
	s.data2[key] = struct{}{}
	return nil
}

func (s Store) unDeclare(key string) {
	delete(s.data2, key)
}

// 变量赋值
func (s *Store) set(key string, value int) error {
	if _, ok := s.data2[key]; ok {
		s.data1[key] = value
		return nil
	}
	return fmt.Errorf("'%s' 没有声明\n", key)
}

func (s *Store) get(key string) (int, error) {
	if v, ok := s.data1[key]; ok {
		return v, nil
	} else if _, ok := s.data2[key]; !ok {
		return -1, fmt.Errorf("变量 '%s' 没有声明", key)
	} else {
		return -1, fmt.Errorf("变量 '%s' 没有初始化", key)
	}
}

type runtime struct {
	data *Store
	err  error
}

func NewRuntime() *runtime {
	return &runtime{
		data: NewStore(),
	}
}

func (r *runtime) evaluate(node *ast.Node) (res *int) {
	var value int
	switch node.Type() {
	case ast.Program:
		for _, child := range node.Children() {
			res = r.evaluate(child)
		}
	case ast.Statement:
		if child := node.ChildrenIndexOf(0); child != nil {
			r.evaluate(child) // 没有返回值
		}
	case ast.IntDeclare:
		// eg: 'int a = 10;' 右结合
		if err := r.data.declare(node.Text()); err != nil {
			fmt.Println(err)
			return
		}
		// 声明并赋值
		if child := node.ChildrenIndexOf(0); child != nil {
			v := r.evaluate(child)
			if v == nil {
				r.data.unDeclare(node.Text())
				return
			}
			value = *v
			if err := r.data.set(node.Text(), value); err != nil {
				r.data.unDeclare(node.Text())
				fmt.Println(err)
				return
			}
		}
	case ast.Assignment:
		if child := node.ChildrenIndexOf(0); child != nil {
			v := r.evaluate(child)
			if v == nil {
				return
			}
			value = *v
			if err := r.data.set(node.Text(), value); err != nil {
				fmt.Println(err)
				return
			}
		}
	case ast.Additive:
		child1 := node.ChildrenIndexOf(0)
		v1 := r.evaluate(child1)

		child2 := node.ChildrenIndexOf(1)
		v2 := r.evaluate(child2)

		if node.Text() == "+" {
			value = *v1 + *v2
		} else {
			value = *v1 - *v2
		}
		res = &value
	case ast.Multiplicative:
		child1 := node.ChildrenIndexOf(0)
		v1 := r.evaluate(child1)

		child2 := node.ChildrenIndexOf(1)
		v2 := r.evaluate(child2)

		if node.Text() == "*" {
			value = *v1 * *v2
		} else if *v2 == 0 {
			panic(fmt.Sprintf("can't division zero\n"))
		} else {
			value = *v1 / *v2
		}
		res = &value
	case ast.Identifier:
		if v, err := r.data.get(node.Text()); err == nil {
			value = v
			res = &value
		} else {
			panic(fmt.Sprintf("%s", err))
		}
	case ast.IntLiteral:
		var err error
		value, err = strconv.Atoi(node.Text())
		if err != nil {
			panic(err)
		}
		res = &value
	}
	return
}
