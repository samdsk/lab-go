//package check
package check

func ResIsPalyn(s string)bool{
  n := len(s)

  if n<=1{
    return true
  }

  if s[0] != s[len(s)-1]{
    return false
  }

  return ResIsPalyn(s[1:len(s)-1])
}
