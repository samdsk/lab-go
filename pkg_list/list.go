package list
import (
  "fmt"
)
type List *Node

type Node struct{
  Content string
  Next *Node
}

func Length(l List) int{
  cursor := l
  c := 0

  for cursor != nil{
    c++
    cursor = cursor.Next
  }

  return c
}

func AddFront(l List, s string) List{
  n := new(Node)
  n.Content = s
  n.Next = l
  return n
}

func AddInOrder(l List, s string) List{

  n := new(Node)
  n.Content = s+"o"

  var prev *Node
  prev = nil
  cursor := l

  for cursor != nil && cursor.Content < s{
    fmt.Println(cursor.Content)
    prev = cursor
    cursor = cursor.Next
  }

  n.Next = cursor

  if prev == nil{
    return n
  }else{
    prev.Next = n
    return l
  }

}

func Concatenate(first, second List)List{
  if first == nil{
    return second
  }

  cursor := first

  for cursor.Next != nil{
    cursor = cursor.Next
    cursor.Next = second
  }

  return first
}

func String(l List)string{
  cursor := l
  t := ""

  for cursor != nil{
    t += cursor.Content + "_"
    cursor = cursor.Next
  }

  return t
}
