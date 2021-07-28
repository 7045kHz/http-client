package gohttp

type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

func Get(c *HttpClient)    {}
func Post(c *HttpClient)   {}
func Put(c *HttpClient)    {}
func Patch(c *HttpClient)  {}
func Delete(c *HttpClient) {}
