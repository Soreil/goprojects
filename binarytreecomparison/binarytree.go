package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
	"sort"
)

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	t1vals := make([]int, 10)
	t2vals := make([]int, 10)

	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		t1vals[i] = <-ch
	}
	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		t2vals[i] = <-ch
	}
	fmt.Println(t1vals)
	fmt.Println(t2vals)
	sort.Ints(t1vals)
	sort.Ints(t2vals)
	strt1 := fmt.Sprintf("%v", t1vals)
	strt2 := fmt.Sprintf("%v", t2vals)
	if strt1 == strt2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
