package ast

type Node struct {
	parent   *Node
	children []*Node
	_type    Type
	_text    string
}

func NewNode(_type Type, _text string) *Node {
	return &Node{
		children: make([]*Node, 0),
		_type:    _type,
		_text:    _text,
	}
}

func (n *Node) Type() Type        { return n._type }
func (n *Node) Text() string      { return n._text }
func (n *Node) Parent() *Node     { return n.parent }
func (n *Node) Children() []*Node { return n.children }

func (n *Node) AddChild(node *Node) {
	// 设置父节点
	node.parent = n
	n.children = append(n.children, node)
}

func (n *Node) ChildrenIndexOf(i int) *Node {
	if i < 0 || i >= len(n.children) {
		return nil
	}
	return n.children[i]
}
