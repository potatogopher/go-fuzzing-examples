package main

import (
  "errors"
  "fmt"
  "unicode/utf8"
)

func main() {
  input := "The quick brown fox jumped over the lazy dog"
  rev, _ := Reverse(input)
  doubleRev, _ := Reverse(rev)
  fmt.Println("original:", input)
  fmt.Println("reversed:", rev)
  fmt.Println("reversed again:", doubleRev)
}

func Reverse(s string) (string, error) {
  if !utf8.ValidString(s) {
    return s, errors.New("input is not valid UTF-8")
  }
  r := []rune(s)
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }
  return string(r), nil
}
