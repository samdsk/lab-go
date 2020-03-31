package gen
import(
  "math/rand"
)
func genChar()rune{
  return 'a'+rune(rand.Intn('z'-'a'))
}
