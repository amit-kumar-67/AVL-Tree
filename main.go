package main

import ("fmt"

)

type Node struct {
	height int
	left *Node
	right *Node
	key int
	
}

func Height(n *Node) int {
	if n==nil {
		return 0
	}
	return n.height

}

func max(a, b int) int{
	if a > b {
     return a
	}
	return b
}

//new Node func
func newNode(key int)*Node {
	node := &Node{}
	node.key =key
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}
//we use here right side rotation
func rightRotate(y *Node) *Node{
	 x := y.left
	 T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(Height(y.left), Height(y.right))+1
	x.height = max(Height(x.left), Height(x.right))+1

	return x
}
//we use here left side roatation
func  leftRotate(x *Node) *Node{
	 y := x.right
	 T2 := y.left

	y.left = x
	x.right =T2

	x.height= max(Height(x.left), Height(x.right))+1
	y.height = max(Height(y.left), Height(y.right))+1

	return y

}
//we use balancing 
func balanceFactor(n *Node)int {
	if n==nil{
		return 0
	}
	return Height(n.left) - Height(n.right)
}
// we use inserting
func insertNode(node *Node, key int) *Node {
	
	if node == nil{
		//return new.node(key)
		return  newNode(key)
	}

	if key < node.key {
		node.left = insertNode(node.left, key)
	}else if key > node.key {
		node.right = insertNode(node.right, key)
	}else {
		return node
	}
	node.height = 1+max(Height(node.left), Height(node.right))

    balance := balanceFactor(node)

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}
	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}
	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}
	 return node
}
func preOdrder(node *Node) *Node{

	pre := node
	for pre.left != nil {
		pre = pre.left
	}
	return pre
	/*
	if node != nil{
		fmt.Println(node.key+" ")
		preOdrder(node.left)
		preOdrder(node.right)
	}*/
}
//we are using for deleting, function is as follows
func deletedNodes(parent *Node, key int) *Node {
	if parent ==nil {
		return parent
	}

	if key < parent.key {
		parent.left = deletedNodes(parent.left, key)
	} else if key > parent.key {
        parent.right = deletedNodes(parent.right, key)
	} else {
		if parent.left == nil || parent.right == nil {
			var temp *Node
			if parent.left ==nil {
				temp = parent.right
			} else {
				temp = parent.left
			}
			if temp == nil {
				temp = parent
				parent =  nil
			} else {
				*parent = *temp
			}
		} else {
			temp := preOdrder(parent.right)
			parent.key = temp.key
			parent.right = deletedNodes(parent.right, temp.key)
		}
	}
	if parent == nil {
		return parent
	}
	//balance the tree
	 parent.height = 1+ max(Height(parent.left), Height(parent.right))
	balance := balanceFactor(parent)
	if balance > 1 {
		if balanceFactor(parent.left) >= 0 {
			return rightRotate(parent)
		} else {
			parent.left = leftRotate(parent.left)
			return rightRotate(parent)
		}
	}
	if balance < -1 {
		if balanceFactor(parent.right) <= 0 {
			return leftRotate(parent)
		} else {
			parent.right = rightRotate(parent.right)
			return leftRotate(parent)
		}
	}
	return parent
}
//we use printing (typethetree)
func typeTheTree(parent *Node, stringtype string, booleantype bool) {
	if parent != nil {
		fmt.Println(stringtype)
		if booleantype {
			fmt.Println("Right side")
			stringtype += " "
		} else {
			fmt.Println("Left side")
			stringtype += "| "
		}
		fmt.Println(parent.key)
		typeTheTree(parent.left, stringtype, false)
		typeTheTree(parent.right, stringtype, true)
	}
}

func main () {
	fmt.Println("Welcome to AVLtree")
	tree := &Node{}

	tree = insertNode(tree, 10)
	tree = insertNode(tree, 20)
	tree= insertNode(tree, 30)
	tree= insertNode(tree, 40)
	tree= insertNode(tree, 50)
	tree= insertNode(tree, 25)
	tree =insertNode(tree, 88)

	typeTheTree(tree, " ", true)
	tree = deletedNodes(tree, 13)
	fmt.Println("when we delete")
	typeTheTree(tree," ", true)
	
	
	
}