package parser

type IteratorResult func() *Result

func (i IteratorResult) ToArray() []*Result {
	arr := make([]*Result, 0, 32)
	for {
		item := i()
		if item == nil {
			break
		}
		arr = append(arr, item)
	}
	return arr
}
