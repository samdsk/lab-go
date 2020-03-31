package gen

func IterPalyn(n int) string{
  t := ""

  if n%2 == 1{
    t = string(genChar())
  }

  for len(t) < n{
    c := genChar()
    t = string(c)+t+string(c)
  }

  return t
}
