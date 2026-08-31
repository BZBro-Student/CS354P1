package customer

// Struct that defines the object customer
type Customer struct {
	name string
}

// Customer constructor
func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c *Customer) String() string {
	return c.name
}
