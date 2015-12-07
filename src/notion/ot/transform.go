package ot

import (
	"fmt"
)

type Transform []interface{}

func (t Transform) Apply(st string) (string, error) {
	nst := st
	cursor := 0
	for _, op := range t {
		switch op.(type) {
		case float64:
			opi := int(op.(float64))
			if opi >= 0 && cursor+opi <= len(nst) {
				cursor += opi
			} else if opi < 0 && cursor - opi <= len(nst) {
				nst = nst[:cursor] + nst[cursor-opi:]
			} else {
				return st, fmt.Errorf("Cursor move operation exceeds length of the string")
			}
		case string:
			nst = nst[:cursor] + op.(string) + nst[cursor:]
			cursor += len(op.(string))
		}
	}
	return nst, nil
}
