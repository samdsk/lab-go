package gen

import(
  "fmt"
)
func ResPalyn(n int)string{
  if n == 0{ return ""}

  if n == 1 {
    return string(genChar())
  }

  t := genChar()

  fmt.Println(t,"updated")

  return string(t)+ResPalyn(n-2)+string(t)
}
