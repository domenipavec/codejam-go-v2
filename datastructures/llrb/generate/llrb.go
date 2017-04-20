package llrb

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "ValueType=int,slice.SliceInt,float64,slice.SliceFloat64,string,slice.SliceString,byte,slice.SliceByte"

// ValueType genny generic type
type ValueType generic.Type

// LLRBValueType tree
type LLRBValueType struct {
	Root *NodeValueType
}

// NodeValueType node
type NodeValueType struct {
	Value ValueType
	Left  *NodeValueType
	Right *NodeValueType
	Red   bool
	Count int
}

// IteratorCallbackValueType iterator
type IteratorCallbackValueType func(v ValueType) bool

// NewValueType construct new llrb
func NewValueType() *LLRBValueType {
	return &LLRBValueType{}
}

// NewNodeValueType construct new node
func NewNodeValueType(v ValueType) *NodeValueType {
	return &NodeValueType{
		Value: v,
		Red:   true,
		Count: 1,
	}
}

// Len returns number of elements in tree
func (llrb LLRBValueType) Len() int {
	return llrb.Root.Len()
}

// Insert an element
func (llrb *LLRBValueType) Insert(v ValueType) {
	n := NewNodeValueType(v)
	llrb.Root = llrb.Root.insert(n)
	llrb.Root.Red = false
}

// Delete an element
func (llrb *LLRBValueType) Delete(v ValueType) {
	llrb.Root = llrb.Root.delete(v)
	llrb.Root.Red = false
}

// Min element
func (llrb LLRBValueType) Min() ValueType {
	min := llrb.Root.Min()
	return min.Value
}

// Max element
func (llrb LLRBValueType) Max() ValueType {
	max := llrb.Root.Max()
	return max.Value
}

// Search if element in tree
func (llrb LLRBValueType) Search(v ValueType) bool {
	return llrb.Root.Search(v)
}

// Get i-th smallest element
func (llrb LLRBValueType) Get(i int) ValueType {
	return llrb.Root.Get(i).Value
}

// Rank of v in tree
func (llrb LLRBValueType) Rank(v ValueType) int {
	return llrb.Root.Rank(v)
}

// IterateAsc iterates ascending
func (llrb LLRBValueType) IterateAsc(cb IteratorCallbackValueType) {
	llrb.Root.IterateAsc(cb)
}

// IterateDesc iterates ascending
func (llrb LLRBValueType) IterateDesc(cb IteratorCallbackValueType) {
	llrb.Root.IterateDesc(cb)
}

// Len returns number of elements under node inclusive
func (node *NodeValueType) Len() int {
	if node == nil {
		return 0
	}
	return node.Count
}

// IsRed checks if node is red
func (node *NodeValueType) IsRed() bool {
	if node == nil {
		return false
	}
	return node.Red
}

// Min returns min element
func (node *NodeValueType) Min() *NodeValueType {
	if node.Left == nil {
		return node
	}
	return node.Left.Min()
}

// Max returns max element
func (node *NodeValueType) Max() *NodeValueType {
	if node.Right == nil {
		return node
	}
	return node.Right.Max()
}

func (node *NodeValueType) refreshCount() {
	node.Count = node.Left.Len() + node.Right.Len() + 1
}

func (node *NodeValueType) rotateLeft() *NodeValueType {
	x := node.Right
	node.Right = x.Left
	x.Left = node
	x.Red = node.Red
	node.Red = true
	node.refreshCount()
	x.refreshCount()
	return x
}

func (node *NodeValueType) rotateRight() *NodeValueType {
	x := node.Left
	node.Left = x.Right
	x.Right = node
	x.Red = node.Red
	node.Red = true
	node.refreshCount()
	x.refreshCount()
	return x
}

func (node *NodeValueType) flipColors() {
	node.Red = !node.Red
	if node.Left != nil {
		node.Left.Red = !node.Left.Red
	}
	if node.Right != nil {
		node.Right.Red = !node.Right.Red
	}
}

func (node *NodeValueType) insert(newNode *NodeValueType) *NodeValueType {
	if node == nil {
		return newNode
	}

	cmp := newNode.Compare(node.Value)
	if cmp == 0 {
		return node
	} else if cmp < 0 {
		node.Left = node.Left.insert(newNode)
	} else {
		node.Right = node.Right.insert(newNode)
	}

	node.refreshCount()

	if node.Right.IsRed() && !node.Left.IsRed() {
		node = node.rotateLeft()
	}
	if node.Left.IsRed() && node.Left.Left.IsRed() {
		node = node.rotateRight()
	}

	if node.Left.IsRed() && node.Right.IsRed() {
		node.flipColors()
	}

	return node
}

func (node *NodeValueType) moveRedLeft() *NodeValueType {
	node.flipColors()
	if node.Right != nil && node.Right.Left.IsRed() {
		node.Right = node.Right.rotateRight()
		node = node.rotateLeft()
		node.flipColors()
	}
	return node
}

func (node *NodeValueType) moveRedRight() *NodeValueType {
	node.flipColors()
	if node.Left != nil && node.Left.Left.IsRed() {
		node = node.rotateRight()
		node.flipColors()
	}
	return node
}

func (node *NodeValueType) delete(v ValueType) *NodeValueType {
	if node.Compare(v) > 0 {
		if node.Left == nil {
			return node
		}
		if !node.Left.IsRed() && !node.Left.Left.IsRed() {
			node = node.moveRedLeft()
		}
		node.Left = node.delete(v)
	} else {
		if node.Left.IsRed() {
			node = node.rotateRight()
		}
		if node.Compare(v) == 0 && node.Right == nil {
			return nil
		}
		if node.Right != nil {
			if !node.Right.IsRed() && !node.Right.Left.IsRed() {
				node = node.moveRedRight()
			}
			if node.Compare(v) == 0 {
				node.Value = node.Right.Min().Value
				node.Right = node.Right.deleteMin()
			} else {
				node.Right = node.Right.delete(v)
			}
		}
	}
	return node.fixUp()
}

func (node *NodeValueType) deleteMin() *NodeValueType {
	if node.Left == nil {
		return nil
	}

	if !node.Left.IsRed() && !node.Left.Left.IsRed() {
		node = node.moveRedLeft()
	}

	node.Left = node.Left.deleteMin()

	return node.fixUp()
}

func (node *NodeValueType) fixUp() *NodeValueType {
	node.refreshCount()

	if node.Right.IsRed() {
		node = node.rotateLeft()
	}

	if node.Left.IsRed() && node.Left.Left.IsRed() {
		node = node.rotateRight()
	}

	if node.Left.IsRed() && node.Right.IsRed() {
		node.flipColors()
	}

	return node
}

// Search if v in tree
func (node *NodeValueType) Search(v ValueType) bool {
	if node == nil {
		return false
	}

	cmp := node.Compare(v)
	if cmp == 0 {
		return true
	}
	if cmp < 0 {
		return node.Left.Search(v)
	}
	return node.Right.Search(v)
}

// Get i-th smallest value
func (node *NodeValueType) Get(i int) *NodeValueType {
	r := node.Left.Len() + 1
	if i == r {
		return node
	}
	if i < r {
		return node.Left.Get(i)
	}
	return node.Right.Get(i - r)
}

// Rank of v in tree
func (node *NodeValueType) Rank(v ValueType) int {
	if node == nil {
		return -1
	}

	r := node.Left.Len() + 1

	cmp := node.Compare(v)
	if cmp == 0 {
		return r
	}
	if cmp < 0 {
		return node.Left.Rank(v)
	}
	return node.Right.Rank(v) + r
}

// IterateAsc iterate ascending
func (node *NodeValueType) IterateAsc(cb IteratorCallbackValueType) bool {
	if node != nil {
		cont := node.Left.IterateAsc(cb)
		if !cont {
			return false
		}
		cont = cb(node.Value)
		if !cont {
			return false
		}
		return node.Right.IterateAsc(cb)
	}
	return true
}

// IterateDesc iterate descending
func (node *NodeValueType) IterateDesc(cb IteratorCallbackValueType) bool {
	if node != nil {
		cont := node.Right.IterateDesc(cb)
		if !cont {
			return false
		}
		cont = cb(node.Value)
		if !cont {
			return false
		}
		return node.Left.IterateDesc(cb)
	}
	return true
}
