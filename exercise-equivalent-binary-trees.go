package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  Walk1(t, ch)
  close(ch)
}

func Walk1(t *tree.Tree, ch chan int) {
  // fmt.Println("Walk1 called")
  if t.Left != nil {
    Walk1(t.Left, ch)
  }
  ch <- t.Value
  if t.Right != nil {
    Walk1(t.Right, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for {
    i, morei := <-ch1
    j, morej := <-ch2
    // fmt.Printf("i = %d, j = %d\n", i, j)
    if i != j {
      return false
    }
    if !morei || !morej {
      if morei != morej {
        return false
      } else {
        return true
      }
    }
  }
}

func main() {
  // Walk the tree
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for {
    i, more := <- ch
    if !more {
      break
    }
    fmt.Println(i)
  }

  // Compare trees
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}

/*
> go run exercise-equivalent-binary-trees.go
1
2
3
4
5
6
7
8
9
10
true
false
*/