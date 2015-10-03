
package util

import (
  "math/rand"
)

const (
  DEFAULT_ID_LENGTH = 10
)

var (
  letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewId() string {
  return NewIdn(DEFAULT_ID_LENGTH)
}

func NewIdn(length int) string {
  b := make([]rune, length)
  for i := range b {
    b[i] = letterRunes[rand.Intn(len(letterRunes))]
  }
  return string(b)
}
