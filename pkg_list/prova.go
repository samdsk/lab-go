package main

import(
  "list"
  "fmt"
  _"os"
)

func main(){

  var l1 list.List

  l1 = list.AddFront(l1,"ciao3")

  l1 = list.AddFront(l1,"ciao2")

  l1 = list.AddFront(l1,"ciao1")

  l1 = list.AddFront(l1,"ciao0")

  l1 = list.AddInOrder(l1,"ciao4")

  l1 = list.AddInOrder(l1,"ciao5")

  l1 = list.AddFront(l1,"ciao6")

  l1 = list.AddInOrder(l1,"ciao7")


  fmt.Println(list.Length(l1))
  fmt.Println(list.String(l1))
}
