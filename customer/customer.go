package customer

// Struct that defines the object customer
type Customer struct {
	name string
}

// Func to construct a customer struct while keeping name private
// from other packages
func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c *Customer) String() string {
	return c.name
}
