package equivalentbinarytree


import "golang.org/x/tour/tree"
import "fmt"
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkChild(t, ch)
	close(ch)
}

func walkChild(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkChild(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walkChild(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch2)

	for {
		i, ok := <- ch
		i2, ok2 := <- ch2
		if ok != ok2 {
			return false
		}
		if !ok {
			break
		}

		if i != i2 {
			return false
		}
	}

	return true
}

func main() {
	result := Same(tree.New(10), tree.New(10))
	fmt.Println(result)
}