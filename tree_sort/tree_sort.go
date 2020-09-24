//Learning URL https://books.studygolang.com/gopl-zh/
package tree_sort

import "math"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) *tree{
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	//return values
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree  {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if t.value < value {
		t.right = add(t.right, value)
	}
	if t.value > value {
		t.left = add(t.left, value)
	}
	return t
}

func KthSmallest(root *tree, k int) int {
	result := int(math.MaxInt32)
	sub(root, &k, &result)
	return result
}

func sub(node *tree, k *int, result *int) {
	if node == nil || *result != int(math.MaxInt32) {
		return
	}

	sub(node.left, k, result)
	if *k == 1 {
		*result = node.value
	}
	*k--
	sub(node.right, k, result)
}