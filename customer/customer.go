package customer

// Struct that defines the object customer
type customer struct {
	name string
}

// Func to construct a customer struct while keeping name private
// from other packages
func Customer(name string) customer {
	return customer{name: name}
}

func (c customer) String() string {
	return c.name
}
