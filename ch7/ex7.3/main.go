package main

// Exercise 7.3: Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4)
// that reveals the sequence of values in the tree.

import (
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

type treeQueue struct {
	queue []*tree
}

func (q *treeQueue) enqueue(t *tree) {
	q.queue = append(q.queue, t)
}

func (q *treeQueue) dequeue() (*tree, error) {
	if len(q.queue) < 1 {
		return nil, fmt.Errorf("can not dequeue an empty queue")
	}
	t := q.queue[0]
	q.queue = q.queue[1:]
	return t, nil
}

func (q *treeQueue) peak() *tree {
	return q.queue[0]
}

func (q *treeQueue) size() int {
	return len(q.queue)
}

func bfs(root *tree, visit func(t *tree, n int)) {
	q := treeQueue{}
	level := 0
	q.enqueue(root)
	q.enqueue(nil)

	for q.size() > 0 {
		node, err := q.dequeue()
		if err != nil {
			return
		}

		if node == nil {
			level++
			q.enqueue(nil)
			if q.peak() == nil {
				break
			}
			continue
		}

		visit(node, level)
		q.enqueue(node.left)
		q.enqueue(node.right)
	}
}

// String Implement BFS to print tree elements
func (root *tree) String() string {
	level := 0
	var sb strings.Builder
	bfs(root, func(t *tree, n int) {
		if n != level {
			sb.WriteString("\n")
			level = n
		}
		sb.WriteString(strconv.Itoa(t.value))
		sb.WriteString(" ")
	})
	return sb.String()
}

func main() {
	t1 := &tree{
		value: 1,
	}

	t2 := &tree{
		value: 2,
	}

	t3 := &tree{
		value: 3,
	}

	t1.left = t2
	t1.right = t3

	t2.left = &tree{
		value: 4,
	}
	t2.right = &tree{
		value: 5,
	}

	t3.left = &tree{
		value: 6,
	}
	t3.right = &tree{
		value: 7,
	}
	fmt.Printf("%v\n", t1)
}
