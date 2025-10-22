package cache

import "fmt"

const Size int = 10

// node definition
type Node struct {
	value string
	left * Node
	right *Node
}

// double linked list
type Queue struct {
	Head * Node
	Tail *Node
	Length int
}

// hashmap
type Hash map[string]*Node

// doubly linked list initializer
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.right = tail
	tail.left = head
	head.value = "head"
	tail.value = "tail"
	return Queue{Head: head, Tail: tail, Length: 0}
}

// cache definition
type Cache struct {
	queue Queue
	hash Hash
}
// Cace initializer
func NewCache() Cache {
	fmt.Println("New cache created successfully")
	return Cache{queue: NewQueue(), hash: Hash{}}
}


// If value exists add it to head. if not still add it to head
func (c * Cache) Check(str string) {
	node := &Node{} // create an empty nor on hrad and store it's address
	// above is similar as malloc() / new node()
	if val , ok := c.hash[str]; ok {
		node = c.Remove(val) 
	} else {
		node = &Node{value: str}
	}
	c.Add(node)
	c.hash[str] = node
}

// remove function

func (c * Cache) Remove(n * Node) *Node {
	fmt.Printf("removing Node with %s", n.value)
	lef := n.left
	righ := n.right
	// from a node we go left and right
	// then we make them point each other
	// the current node goes out of scope and is clean my GC
	righ.left = lef
	lef.right = righ
	c.queue.Length--
	delete(c.hash, n.value) // remove from the map also
	// this is done so we don't point to old node cuz we create a new node and need to add it to the map
	return n // return the node that has now been modified  
}

// add to cache

func (c * Cache) Add(n * Node) {
	fmt.Printf("Adding Node to the cache with value: %s",n.value)
	temp := c.queue.Head.right // this is same as Node * temp = head->right
	c.queue.Head.right = n; // now head points to the new head ADD(n)
	n.left= c.queue.Head // this left now points to head whose left is nullptr.
	n.right = temp // the n.left points to the previus head which is  now the second element. essentially the next of head
	c.queue.Length++

	if c.queue.Length > Size {
		c.Remove(c.queue.Tail.left) // remove the least used element
	}
}

func (q * Queue) Display() {
	tempNode := q.Head.right // get the first node

	fmt.Printf("%d - [", q.Length)

	for i:=0; i < q.Length; i++ {
		fmt.Printf("{%s}", tempNode.value)
		if  i < q.Length -1 {
			fmt.Printf("--->")
		}
		tempNode = tempNode.right
	}
	fmt.Println("]")
}
func (c * Cache)Display() {
	c.queue.Display()
}