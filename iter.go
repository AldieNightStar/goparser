package parser

type IteratorResult func() (*Result, int)

func (i IteratorResult) ToArray() ([]*Result, int) {
	arr := make([]*Result, 0, 32)
	pos := 0
	for {
		item, newPos := i()
		if item == nil {
			break
		}
		pos = newPos
		arr = append(arr, item)
	}
	return arr, pos
}

func (i IteratorResult) FilterArray(filter func(*Result) bool) ([]*Result, int) {
	arr := make([]*Result, 0, 32)
	pos := 0
	for {
		item, p := i()
		if item == nil {
			break
		}
		ok := filter(item)
		if !ok {
			continue
		}
		pos = p
		arr = append(arr, item)
	}
	return arr, pos
}

func (i IteratorResult) UntilArray(until func(*Result) bool) ([]*Result, int) {
	arr := make([]*Result, 0, 32)
	pos := 0
	for {
		item, p := i()
		isEnd := until(item)
		if isEnd || item == nil {
			break
		}
		pos = p
		arr = append(arr, item)
	}
	return arr, pos
}

func (i IteratorResult) FewArray(cnt int) ([]*Result, int) {
	arr := make([]*Result, 0, 32)
	pos := 0
	for {
		item, p := i()
		if len(arr) >= cnt {
			break
		}
		if item == nil {
			break
		}
		pos = p
		arr = append(arr, item)
	}
	return arr, pos
}
