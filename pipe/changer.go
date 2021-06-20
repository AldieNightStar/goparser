package pipe

type Changer struct {
	name string
	list []interface{}
	cb   func(*ChangerValue)
}

type ChangerValue struct {
	Name string
	List []interface{}
}

func NewChanger(cb func(*ChangerValue)) *Changer {
	c := &Changer{}
	c.list = make([]interface{}, 0, 32)
	c.cb = cb
	return c
}

func NewChangerToList(list *[]*ChangerValue) *Changer {
	return NewChanger(func(cv *ChangerValue) {
		*list = append(*list, cv)
	})
}

func (c *Changer) Put(name string, val interface{}) {
	if c.name == "" {
		c.name = name
	}
	if c.name == name {
		c.list = append(c.list, val)
	} else {
		c.cb(&ChangerValue{c.name, c.list})
		c.name = name
		c.list = make([]interface{}, 0, 32)
		c.list = append(c.list, val)
	}
}

func (c *Changer) Done() {
	if len(c.list) > 0 {
		c.cb(&ChangerValue{c.name, c.list})
		c.name = ""
		c.list = make([]interface{}, 0, 32)
	}
}
