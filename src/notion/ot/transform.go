package ot

type Transform []interface{}

func (t Transform) Apply(st string) string {
	cursor := 0
	for _, op := range t {
		switch op.(type) {
		case int:
			if op.(int) >= 0 {
				cursor += op.(int)
			} else {
				st = st[:cursor+op.(int)] + st[cursor+1:]
			}
		case string:
			st = st[:cursor] + op.(string) + st[cursor+len(op.(string))+1:]
		}
	}
	return st
}
