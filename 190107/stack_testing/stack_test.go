package main

import(
  "testing"
  "fmt"

)


func TestPush(t *testing.T){
  fmt.Println("Testing push function")
  i := 1.
  push(i)
  n := stack[0]

  if n != 1{
    t.Errorf("Fail: Expected: %d, Output: %d\n",int(i),int(n))
  }
}
