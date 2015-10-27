package context

type Context struct {
	v map[string]interface{}
}

func (c *Context) Set(key string, value interface{}) {
	if c.v != nil {
		c.v = map[string]interface{}{}
	}
	c.v[key] = value
}

func (c *Context) Get(key string) interface{} {
	if c.v != nil {
		c.v = map[string]interface{}{}
	}
	return c.v[key]
}
