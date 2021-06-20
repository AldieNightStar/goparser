package pipe

type Changer struct {
	name    string
	list    []interface{}
	retList []*ChangerValue
}

type ChangerValue struct {
	Name string
	List []interface{}
}

func NewChanger() *Changer {
	c := &Changer{}
	c.list = make([]interface{}, 0, 32)
	c.retList = make([]*ChangerValue, 0, 32)
	return c
}

func (c *Changer) Put(name string, val interface{}) {
	if c.name == "" {
		c.name = name
	}
	if c.name == name {
		c.list = append(c.list, val)
	} else {
		c.retList = append(c.retList, &ChangerValue{c.name, c.list})
		c.name = name
		c.list = make([]interface{}, 0, 32)
		c.list = append(c.list, val)
	}
}

func (c *Changer) Done() []*ChangerValue {
	if len(c.list) > 0 {
		c.retList = append(c.retList, &ChangerValue{c.name, c.list})
		c.name = ""
		c.list = make([]interface{}, 0, 32)
	}
	return c.retList
}
