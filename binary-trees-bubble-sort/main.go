package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value // Add value to channel

	// If both branches exist, search both
	if t.Left != nil && t.Right != nil {
		go Walk(t.Left, ch)
		go Walk(t.Right, ch)
		return
	}

	// If there's a left branch, search left
	if t.Left != nil {
		go Walk(t.Left, ch)
		return
	}

	// If there's a right branch, search right
	if t.Right != nil {
		go Walk(t.Right, ch)
		return
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	ls1, ls2 := [10]int{}, [10]int{}

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v, ok := <-ch1
		if !ok {
			break
		}

		ls1[i] = v
	}

	for i := 0; i < 10; i++ {
		v, ok := <-ch2
		if !ok {
			break
		}

		ls2[i] = v
	}

	// Sort list's
	ls1 = BubbleSort(ls1)
	ls2 = BubbleSort(ls2)

	return ls1 == ls2
}

func BubbleSort(ls [10]int) [10]int {
    for {
        var sw bool // If this variable is false, sorting is done
        for i := 0; i < len(ls) - 1; i++ {
            if ls[i] > ls[i+1] {
                ls[i], ls[i+1] = ls[i+1], ls[i] // Set previous element to next
                sw = true // Set sw var to true so program continues
            }
        }
        if !sw { break } // End loop when no more swaps have been executed
    }
	return ls
}

func main() {
    fmt.Println("Testing walk function...")
	ch := make(chan int)
	t := tree.New(1)

	go Walk(t, ch)

	// Print results from func call above
	for i := 0; i < 10; i++ {
		v, ok := <-ch
		if !ok {
			break
		}

		fmt.Println(v)
	}

	// Close channel
	close(ch)
    fmt.Println("Done!")

    fmt.Println("Testing same function...")
	// Check if trees are equal
	fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
    fmt.Println("Done!")
}
