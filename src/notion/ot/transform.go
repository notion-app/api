package ot

import (
	"fmt"
)

type Transform struct {
	Type    string `json:"type"`
	At      int    `json:"at"`
	Content string `json:"content"`
}

func (t Transform) Do(st string) (string, error) {
	switch t.Type {
	case "insert":
		return st[:t.At] + t.Content + st[t.At:], nil
	case "delete":
		return st[:t.At] + st[t.At+1:], nil
	}
	return st, fmt.Errorf("Transform type must be either 'insert' or 'delete'")
}
