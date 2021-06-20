package pipe

import "testing"

func TestChanger(t *testing.T) {
	list := make([]*ChangerValue, 0, 32)
	c := NewChangerToList(&list)

	c.Put("A", 32)
	c.Put("A", 64)
	c.Put("A", 48)

	c.Put("B", 12)
	c.Put("B", 24)

	c.Put("C", 33)
	c.Put("C", 4)

	c.Done()

	c1 := list[0]
	c2 := list[1]
	c3 := list[2]

	if c1.Name != "A" {
		t.Fatal("Channel first value is not A")
	}
	if len(c1.List) != 3 {
		t.Fatal("List 1 of elems not 3. ", len(c1.List))
	}

	if c2.Name != "B" {
		t.Fatal("Channel first value is not B")
	}
	if len(c2.List) != 2 {
		t.Fatal("List 2 of elems not 2. ", len(c2.List))
	}

	if c3.Name != "C" {
		t.Fatal("Channel first value is not C")
	}
	if len(c3.List) != 2 {
		t.Fatal("List 3 of elems not 2. ", len(c3.List))
	}

}
