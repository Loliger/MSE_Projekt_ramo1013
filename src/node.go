package src

import "fmt"

type node[T any] struct {
	val  T
	next *node[T]
}

func mkNode[T any](v T) *node[T] {
	return &node[T]{val: v, next: nil}
}

func insert[T any](n *node[T], v T) {
	n.next = mkNode[T](v)
}

// n > 0
func mkList[T any](n int, v T) *node[T] {

	head := mkNode(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert(current, v)
		current = current.next

	}
	return head

}

func len[T any](n *node[T]) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

func TestNode() {
	n1 := mkList[int](10, 1)
	fmt.Printf("\n%d", len(n1))
	fmt.Printf("\n%T", n1.val)

	n2 := mkList[bool](10, true)
	fmt.Printf("\n%d", len(n2))
	fmt.Printf("\n%T", n2.val)
}

// Generic translation (sketch)
type node_G struct {
	val  interface{}
	next *node_G
}

func mkNode_G(v interface{}) *node_G {
	return &node_G{val: v, next: nil}
}

func insert_G(n *node_G, v interface{}) {
	n.next = mkNode_G(v)
}

// n > 0
func mkList_G(n int, v interface{}) *node_G {

	head := mkNode_G(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_G(current, v)
		current = current.next

	}
	return head
}

func len_G(n *node_G) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}
	return i
}

func TestNode_G() {
	n1 := mkList_G(10, 1)
	fmt.Printf("\n%d", len_G(n1))
	fmt.Printf("\n%T", n1.val)

	n2 := mkList_G(10, true)
	fmt.Printf("\n%d", len_G(n2))
	fmt.Printf("\n%T", n2.val)
}

// Monomorphization (sketch)
type node_int struct {
	val  int
	next *node_int
}

func mkNode_int(v int) *node_int {
	return &node_int{val: v, next: nil}
}

func insert_int(n *node_int, v int) {
	n.next = mkNode_int(v)
}

// n > 0
func mkList_int(n int, v int) *node_int {

	head := mkNode_int(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_int(current, v)
		current = current.next

	}
	return head
}

func len_int(n *node_int) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}
	return i
}

type node_bool struct {
	val  bool
	next *node_bool
}

func mkNode_bool(v bool) *node_bool {
	return &node_bool{val: v, next: nil}
}

func insert_bool(n *node_bool, v bool) {
	n.next = mkNode_bool(v)
}

// n > 0
func mkList_bool(n int, v bool) *node_bool {

	head := mkNode_bool(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_bool(current, v)
		current = current.next

	}
	return head
}

func len_bool(n *node_bool) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}
	return i
}

func TestNode_MM() {
	n1 := mkList_int(10, 1)
	fmt.Printf("\n%d", len_int(n1))
	fmt.Printf("\n%T", n1.val)

	n2 := mkList_bool(10, true)
	fmt.Printf("\n%d", len_bool(n2))
	fmt.Printf("\n%T", n2.val)
}
