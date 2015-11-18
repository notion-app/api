
package ot

import (
  "testing"
)

func TestTransformApply(t *testing.T) {
  xf1 := Transform{1,2,3}
  r1 := xf1.Apply("hello")
  if r1 != "hello" {
    t.Error("XF1 should not modify the string")
  }
  xf2 := Transform{1,-1}
  r2 := xf2.Apply("hello")
  if r2 != "ello" {
    t.Error("XF2 should only remove the first character")
  }
  xf3 := Transform{2,"xx"}
  r3 := xf3.Apply("hello")
  if r3 != "hexxllo" {
    t.Error("XF3 should add an 'xx' two characters into the string")
  }
  xf4 := Transform{
    1,
    -1,
    2,
    "zz",
    -1,
    "xz",
  }
  r4 := xf4.Apply("hello")
  if r4 != "elzxzlo" {
    t.Error("XF4 did not return the proper response")
  }
}
