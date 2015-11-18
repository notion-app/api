package ot

type Transform []interface{}

func (t Transform) Apply(st string) string {
	cursor := 0
	for _, op := range t {
		switch op.(type) {
		case int:
			if op.(int) >= 0 && cursor + op.(int) < len(st) {
				cursor += op.(int)
			} else if op.(int) < 0 && cursor + op.(int) >= 0 {
				st = st[:cursor+op.(int)] + st[cursor:]
				cursor += op.(int)
			}
		case string:
			st = st[:cursor] + op.(string) + st[cursor:]
			cursor += len(op.(string))
		}
	}
	return st
}
