package parser

type IteratorResult func() (*Result, int)

func (i IteratorResult) ToArray() []*Result {
	arr := make([]*Result, 0, 32)
	for {
		item, _ := i()
		if item == nil {
			break
		}
		arr = append(arr, item)
	}
	return arr
}

func (i IteratorResult) FilterArray(filter func(*Result) bool) []*Result {
	arr := make([]*Result, 0, 32)
	for {
		item, _ := i()
		if item == nil {
			break
		}
		ok := filter(item)
		if !ok {
			continue
		}
		arr = append(arr, item)
	}
	return arr
}

func (i IteratorResult) UntilArray(until func(*Result) bool) []*Result {
	arr := make([]*Result, 0, 32)
	for {
		item, _ := i()
		isEnd := until(item)
		if isEnd || item == nil {
			break
		}
		arr = append(arr, item)
	}
	return arr
}

func (i IteratorResult) FewArray(cnt int) []*Result {
	arr := make([]*Result, 0, 32)
	for {
		item, _ := i()
		if len(arr) >= cnt {
			break
		}
		if item == nil {
			break
		}
		arr = append(arr, item)
	}
	return arr
}
