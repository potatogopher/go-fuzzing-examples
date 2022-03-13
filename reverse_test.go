package main

import (
  "testing"
  "unicode/utf8"
)

// Old Test --
// Limitations are that we need to manually come up with all the
// inputs that we want to test.
// func TestReverse(t *testing.T) {
//   for name, test := range map[string]struct{
//     in, want string
//   }{
//     "test 1": {"Hello, world", "dlrow ,olleH"},
//     "test 2": {" ", " "},
//     "test 3": {"12345!", "!54321"},
//   }{
//     t.Run(name, func(t *testing.T) {
//       rev := Reverse(test.in)
//       if rev != test.want {
//         t.Errorf("Reverse: %q, want %q", rev, test.want)
//       }
//     })
//   }
// }

// FuzzReverse tests the reverse function while coming up with new inputs
// and may identify edge cases.
//
// Limitations: We can't predict expected outputs since we don't have
// control over the inputs. However, we can use the Reverse function to
// do some verification within out test.
//
// 1. Reversing the string twice results in the original value.
// 2. The reversed string is valid UTF-8.
//
// If invalid UTF-8 values are in the string, we will skip the test.
func FuzzReverse(f *testing.F) {
  testcases := []string{"Hello, world!", " ", "12345!"}
  // add cases to corpus
  for _, tc := range testcases {
    f.Add(tc)
  }
  f.Fuzz(func(t *testing.T, orig string) {
    rev, err := Reverse(orig)
    if err != nil {
      return
    }
    doubleRev, err := Reverse(rev)
    if err != nil {
      return
    }
    if orig != doubleRev {
      t.Errorf("Before: %q, after %q", orig, rev)
    }
    if utf8.ValidString(orig) && !utf8.ValidString(rev) {
      t.Errorf("Reverse procuded invalid UTF-8 string %q", rev)
    }
  })
}
